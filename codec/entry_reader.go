package codec

import (
	"bytes"
	"io"
)

type EntryContentReader func(dict *Dict, idx *DictEntry, typ ContentType, buf *bytes.Buffer) (*EntryContent, error)

var EntryContentReaders = make(map[ContentType]EntryContentReader)

type EntryContent struct {
	Type ContentType
	Text string
}

func readTextEntry(dict *Dict, idx *DictEntry, typ ContentType, buf *bytes.Buffer) (*EntryContent, error) {
	entry := &EntryContent{}
	entry.Type = typ
	b, err := buf.ReadBytes(0)
	if err != nil && err != io.EOF {
		return nil, err
	}
	if len(b) > 0 && b[len(b)-1] == 0 {
		entry.Text = string(b[:len(b)-1])
	} else {
		entry.Text = string(b)
	}
	return entry, nil
}

func init() {
	EntryContentReaders[NULL_TERMINAL_TEXT] = readTextEntry
	EntryContentReaders[ENGLISH_PHONETIC] = readTextEntry
	EntryContentReaders[YINBIAO] = readTextEntry
	EntryContentReaders[KINGSOFT_XML] = readTextEntry
	EntryContentReaders[HTML] = readTextEntry
	EntryContentReaders[XDXF_MARKUP] = readTextEntry
	EntryContentReaders[PANGO_TEXT] = readTextEntry
}
