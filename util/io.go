package stardictutil

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/wenerme/letsgo/compress"
	"github.com/wenerme/stardict/db"
	"github.com/wenerme/stardict/fmt"
	"github.com/wenerme/stardict/genproto/v1/stardictdata"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Open and convert to data
func Read(file string) (data *stardictdata.StardictData, err error) {
	return ReadWithFormat(file, "")
}
func ReadWithFormat(file string, format string) (data *stardictdata.StardictData, err error) {
	if format == "" {
		switch filepath.Ext(wcompress.FinalName(file)) {
		case ".json":
			format = "json"
		case ".pb":
			format = "pb"
		case ".ifo", ".idx", ".dict", ".syn", ".tar":
			format = "raw"
		default:
			err = errors.New("can not detect format from filename")
			return
		}
	}
	switch format {

	case "raw":
		var dict *stardictfmt.Reader
		dict, err = stardictfmt.Read(file)
		if err != nil {
			return
		}
		data, err = ConvertRawToData(dict)
	case "json", "pb":
		var b []byte
		if _, b, err = wcompress.DecompressAll(file); err != nil {
			return
		}
		data = &stardictdata.StardictData{}
		switch format {
		case "pb":
			err = proto.Unmarshal(b, data)
		case "json":
			err = json.Unmarshal(b, data)
		}
	default:
		err = errors.New("invalid file format")
	}
	return
}

// Write date to
func Write(data *stardictdata.StardictData, path string) (err error) {
	return WriteWithFormat(data, path, "")
}
func WriteWithFormat(data *stardictdata.StardictData, file string, format string) (err error) {
	if format == "" {
		switch filepath.Ext(file) {
		case ".pb":
			format = "protobuf"
		case ".db":
			format = "sqlite"
		case ".csv":
			format = "csv"
		case ".json":
			format = "json"
		default:
			err = errors.New("can not detect format from filename")
			return
		}
	}
	var b []byte
	switch format {
	case "csv":
	case "sqlite":
		return writeDb(data, file, "sqlite3")
	case "json":
		b, err = json.Marshal(data)
		if err != nil {
			return
		}
		err = ioutil.WriteFile(file, b, os.ModePerm)
	case "pb":
		fallthrough
	case "protobuf":
		b, err = proto.Marshal(data)
		if err != nil {
			return
		}
		err = ioutil.WriteFile(file, b, os.ModePerm)
	default:
		err = errors.New("invalid format")
	}
	return
}
func writeDb(data *stardictdata.StardictData, dialect string, arg string) (err error) {
	var db *gorm.DB
	if db, err = gorm.Open(dialect, arg); err != nil {
		return
	}
	defer db.Close()
	if err = db.AutoMigrate(
		stardictdb.Dictionary{},
		stardictdb.Word{},
		stardictdb.Synonym{},
		stardictdb.Content{}).Error; err != nil {
		return
	}

	{
		db := db.Begin()
		defer func() {
			if err != nil {
				db.Rollback()
			} else {
				db.Commit()
			}
		}()
		info := data.Info
		dict := &stardictdb.Dictionary{
			Code:         info.Code,
			Name:         info.Name,
			Version:      info.Version,
			Description:  info.Description,
			Author:       info.Author,
			Email:        info.Email,
			Website:      info.Website,
			Type:         info.Type,
			WordCount:    int(info.WordCount),
			SynonymCount: int(info.SynonymCount),
		}
		if info.Date != nil {
			t := time.Unix(int64(info.Date.GetSeconds()), 0)
			dict.Date = &t
		}
		if err = db.Create(dict).Error; err != nil {
			err = errors.Wrap(err, "failed to insert dict")
			return
		}

		for _, v := range data.Entries {
			word := &stardictdb.Word{
				Dict: dict,
				Word: v.Word,
			}
			logrus := logrus.WithField("word", word.Word).WithField("dict", dict.ID)
			if err = db.Create(word).Error; err != nil {
				err = errors.Wrap(err, "failed to insert word")
				logrus.WithError(err).Warn("failed to insert word")
				return
			}
			for _, c := range v.Contents {
				content := &stardictdb.Content{
					Dict:    dict,
					Word:    word,
					Type:    c.Type.String(),
					Content: c.Text,
				}
				if err = db.Create(content).Error; err != nil {
					err = errors.Wrap(err, "failed to insert content")
					logrus.WithField("type", c.Type.String()).Warn("failed to insert content")
					return
				}
			}
			for _, s := range v.Synonyms {
				synonym := &stardictdb.Synonym{
					Dict:    dict,
					Word:    word,
					Synonym: s,
				}
				if err = db.Create(synonym).Error; err != nil {
					err = errors.Wrap(err, "failed to insert synonym")
					logrus.WithError(err).WithField("synonym", s).Warn("failed to insert synonym")
					return
				}
			}
		}
	}
	return
}
