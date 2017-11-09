package stardictutil

import (
	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/wenerme/letsgo/compress"
	"github.com/wenerme/stardict/codec"
	"github.com/wenerme/stardict/db"
	"github.com/wenerme/stardict/genproto/v1/stardictdata"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Open and convert to data
func Open(file string) (data *stardictdata.StardictData, err error) {
	switch filepath.Ext(wcompress.FinalName(file)) {
	case ".pb":
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		_, r, err := wcompress.Decompress(f, file)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}
		data = &stardictdata.StardictData{}
		err = proto.Unmarshal(b, data)
	case ".ifo", ".idx", ".dict", ".syn", ".tar":
		var dict *codec.Dict
		dict, err = codec.Open(file)
		if err != nil {
			return
		}
		data, err = ConvertRawToData(dict)
	default:
		err = errors.New("invalid file format")
	}
	return
}

// Write date to
func Write(data *stardictdata.StardictData, path string, format string) (err error) {
	var db *gorm.DB
	if db, err = gorm.Open("sqlite3", path); err != nil {
		return
	}
	defer db.Close()
	db.AutoMigrate(stardictdb.Dictionary{}, stardictdb.Word{}, stardictdb.Synonym{}, stardictdb.Content{})

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
