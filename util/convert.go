package stardictutil

import (
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
	"github.com/wenerme/stardict/fmt"
	"github.com/wenerme/stardict/genproto/v1/stardictdata"
)

func ConvertRawToData(dict *stardictfmt.Reader) (*stardictdata.StardictData, error) {
	data := &stardictdata.StardictData{}
	data.Info = convertInfo(dict.Info)

	for _, v := range dict.Entries {
		entry, err := convertEntry(v)
		if err != nil {
			return nil, err
		}
		data.Entries = append(data.Entries, entry)
	}
	return data, nil
}

func convertEntry(entry *stardictfmt.DictEntry) (*stardictdata.Entry, error) {
	data := &stardictdata.Entry{}
	data.Word = entry.Word
	data.Synonyms = entry.Synonyms
	for _, v := range entry.Contents {
		content, err := convertContent(v)
		if err != nil {
			return nil, err
		}
		data.Contents = append(data.Contents, content)
	}

	return data, nil
}

func convertContent(raw *stardictfmt.EntryContent) (data *stardictdata.Content, err error) {
	data = &stardictdata.Content{}
	data.Text = raw.Text
	data.Type, err = convertTypeRawToData(raw.Type)
	return
}
func convertTypeRawToData(raw stardictfmt.ContentType) (stardictdata.ContentType, error) {
	switch raw {
	case stardictfmt.NULL_TERMINAL_TEXT:
		return stardictdata.ContentType_TEXT, nil
	case stardictfmt.HTML:
		return stardictdata.ContentType_HTML, nil
	case stardictfmt.YINBIAO:
		return stardictdata.ContentType_YIBIAO, nil
	case stardictfmt.ENGLISH_PHONETIC:
		return stardictdata.ContentType_PHONETI, nil
	case stardictfmt.XDXF_MARKUP:
		return stardictdata.ContentType_PHONETI, nil
	}
	return 0, errors.New("unknown content type raw to data: " + string(raw))
}

func convertInfo(info *stardictfmt.DictInfo) *stardictdata.Info {
	data := &stardictdata.Info{}
	data.Name = info.BookName
	data.Version = info.Version
	data.Description = info.Description
	data.Type = info.DictType
	if info.Date.Second() != 0 {
		data.Date = new(google_protobuf.Timestamp)
		data.Date.Seconds = int64(info.Date.Second())
	}
	data.Author = info.Author
	data.Email = info.Email
	data.Website = info.Website

	data.WordCount = int32(info.WordCount)
	data.SynonymCount = int32(info.SynWordCount)
	return data
}
