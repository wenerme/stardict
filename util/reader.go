package stardictutil

import (
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/wenerme/letsgo/compress"
	"github.com/wenerme/stardict/codec"
	"github.com/wenerme/stardict/genproto/v1/stardictdata"
	"io/ioutil"
	"os"
	"path/filepath"
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
	case ".ifo", ".idx", ".dict", ".syn":
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
