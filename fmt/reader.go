package stardictfmt

import (
	"archive/tar"
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
	"github.com/wenerme/letsgo/compress"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type LoadConf struct {
	ArchiveFile   string
	ArchiveReader io.Reader

	InfoFile  string
	IndexFile string
	DictFile  string
	SynFile   string

	InfoReader  io.Reader
	IndexReader io.Reader
	DictReader  io.Reader
	SynReader   io.Reader

	closers []io.Closer
}

func (self *LoadConf) Info() (r io.Reader, err error) {
	if self.InfoReader != nil {
		return self.InfoReader, nil
	}
	if self.InfoFile != "" {
		return os.Open(self.InfoFile)
	}
	return nil, errors.New("info not found")
}
func (self *LoadConf) Index() (r io.Reader, err error) {
	if self.IndexReader != nil {
		return self.IndexReader, nil
	}
	if self.IndexFile != "" {
		return os.Open(self.IndexFile)
	}
	return nil, os.ErrNotExist
}
func (self *LoadConf) Dict() (r io.Reader, err error) {
	if self.DictReader != nil {
		return self.DictReader, nil
	}
	if self.DictFile != "" {
		return self.Open(self.DictFile)
	}
	return nil, os.ErrNotExist
}
func (self *LoadConf) Syn() (io.Reader, error) {
	if self.SynReader != nil {
		return self.SynReader, nil
	}
	if self.SynFile != "" {
		return self.Open(self.SynFile)
	}
	return nil, os.ErrNotExist
}
func (self *LoadConf) Open(file string) (r io.Reader, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	self.AppendCloser(f)
	_, r, err = wcompress.Decompress(f, file)
	return
}
func (self *LoadConf) AppendCloser(c io.Closer) {
	self.closers = append(self.closers, c)
}

func (self *LoadConf) Close() (err error) {
	for _, c := range self.closers {
		e := c.Close()
		if e != nil {
			if err == nil {
				e = err
			} else {
				err = errors.Wrap(e, err.Error())
			}
		}
	}
	return
}

func Read(file string) (dict *Reader, err error) {
	var conf *LoadConf
	conf, err = NewConf(file)
	if err != nil {
		return
	}
	defer conf.Close()
	d := NewReader()
	err = d.LoadWithConf(conf)
	if err != nil {
		return
	}
	dict = d
	return
}
func NewConf(file string) (conf *LoadConf, err error) {
	ext := filepath.Ext(wcompress.FinalName(file))
	switch ext {
	case ".tar":
		return newArchiveConf(file)
	case ".ifo", ".idx", ".dict", ".syn":
		return newFileConf(file[:len(file)-len(ext)], "")
	default:
		err = errors.New("unknown file format")
	}
	return
}

func newFileConf(dir string, fn string) (conf *LoadConf, err error) {
	base := filepath.Join(dir, fn)
	conf = &LoadConf{}
	conf.InfoFile = findFile(base, ".ifo")
	conf.IndexFile = findFile(base, ".idx")
	conf.DictFile = findFile(base, ".dict", ".dict.dz")
	conf.SynFile = findFile(base, ".syn")
	if conf.InfoFile == "" {
		err = errors.New("no .ifo file")
	}
	if conf.IndexFile == "" {
		err = errors.New("no .idx file")
	}
	if conf.DictFile == "" {
		err = errors.New("no dict file")
	}
	return
}
func newArchiveConf(file string) (conf *LoadConf, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	defer f.Close()
	fn, r, err := wcompress.Decompress(f, file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decompress")
	}

	conf = &LoadConf{}
	defer func() {
		if err != nil {
			conf.Close()
			conf = nil
		}
	}()
	switch filepath.Ext(fn) {
	case ".tar":
		reader := tar.NewReader(r)
		var header *tar.Header
		for {
			header, err = reader.Next()
			if err == io.EOF {
				err = nil
				break
			} else if err != nil {
				return
			}

			b := make([]byte, header.Size)
			_, err = io.ReadFull(reader, b)
			if err != nil {
				return
			}
			name := header.Name
			name, r, err = wcompress.Decompress(bytes.NewReader(b), name)
			if err != nil {
				return
			}
			switch filepath.Ext(name) {
			case ".idx":
				conf.IndexReader = r
			case ".ifo":
				conf.InfoReader = r
			case ".dict":
				conf.DictReader = r
			case ".syn":
				conf.SynReader = r
			}
		}
	default:
		err = errors.New("unknown archive file: " + fn)
	}
	return
}

func NewReader() *Reader {
	return &Reader{}
}

func (self *Reader) LoadWithConf(conf *LoadConf) (err error) {
	var r io.Reader
	if r, err = conf.Info(); err != nil {
		return
	} else if err = self.LoadInfo(r); err != nil {
		return
	}

	if r, err = conf.Index(); err != nil {
		return
	} else if err = self.LoadIndex(r); err != nil {
		return
	}

	if r, err = conf.Dict(); err != nil {
		return
	} else if err = self.LoadDict(r); err != nil {
		return
	}

	if r, err = conf.Syn(); err != nil {
		if err == os.ErrNotExist {
			err = nil
		} else {
			return
		}
	} else if r != nil {
		err = self.LoadSynonym(r)
	}

	return
}

func (self *Reader) LoadInfoFile(file string) (err error) {
	f, err := openFile(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return self.LoadInfo(f)
}
func (self *Reader) LoadIndexFile(file string) (err error) {
	f, err := openFile(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return self.LoadIndex(f)
}
func (self *Reader) LoadInfo(r io.Reader) (err error) {
	self.Info, err = readInfo(r)
	return
}
func (self *Reader) LoadIndex(r io.Reader) (err error) {
	if self.Info == nil {
		return errors.New("no info loaded yet")
	}
	self.Entries, err = readIndexEntries(self.Info.IndexOffsetBits, r)
	return
}
func (self *Reader) LoadDictFile(file string) (err error) {
	f, err := openFile(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return self.LoadDict(f)
}
func (self *Reader) LoadDict(r io.Reader) (err error) {
	for _, idx := range self.Entries {
		err = readDictContent(self, idx, r)
		if err != nil {
			return
		}
	}
	return nil
}

func readInfo(reader io.Reader) (info *DictInfo, err error) {
	info = &DictInfo{}
	scanner := bufio.NewScanner(reader)

	if !scanner.Scan() {
		return nil, errors.New("info magic header not found")
	}
	info.Header = scanner.Text()

	for scanner.Scan() {
		err = scanner.Err()
		if err == io.EOF {
			err = nil
			break
		}

		if err != nil {
			return nil, err
		}
		line := scanner.Text()
		i := strings.IndexRune(line, '=')
		if i < 1 {
			return nil, errors.New("invalid info line: " + line)
		}

		key := strings.TrimSpace(line[:i])
		val := strings.TrimSpace(line[i+1:])
		switch key {
		case "version":
			info.Version = val
		case "wordcount":
			info.WordCount, err = strconv.Atoi(val)
		case "synwordcount":
			info.SynWordCount, err = strconv.Atoi(val)
		case "idxfilesize":
			info.IndexFileSize, err = strconv.Atoi(val)
		case "idxoffsetbits":
			info.IndexOffsetBits, err = strconv.Atoi(val)
		case "bookname":
			info.BookName = val
		case "description":
			info.Description = val
		case "date":
			info.Date, err = parseDate(val)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to parse date '%v'", val)
				err = nil
			}
		case "sametypesequence":
			info.SameTypeSequence = val
		case "dicttype":
			info.DictType = val
		case "author":
			info.Author = val
		case "email":
			info.Email = val
		case "website":
			info.Website = val
		default:
			return nil, errors.New("unknown info line: " + line)
		}
		if err != nil {
			return nil, errors.New("failed to parse info line: " + line)
		}
	}
	if info.IndexOffsetBits == 0 {
		info.IndexOffsetBits = 32
	}
	return
}

func readIndexEntries(bits int, reader io.Reader) (entries []*DictEntry, err error) {
	r := bufio.NewReader(reader)
	buf := make([]byte, 4)
	var b []byte
	for {
		entry := &DictEntry{}

		// word
		b, err = r.ReadBytes(0)
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return nil, err
		}
		entry.Word = string(b[:len(b)-1]) // trim \x00

		// offset
		if bits == 32 {
			_, err = io.ReadFull(r, buf[:4])
			entry.Offset = uint64(binary.BigEndian.Uint32(buf))
		} else {
			_, err = io.ReadFull(r, buf)
			entry.Offset = binary.BigEndian.Uint64(buf)
		}
		if err != nil {
			return nil, err
		}

		// size
		_, err = io.ReadFull(r, buf[:4])
		if err != nil {
			return nil, err
		}
		entry.Size = int(binary.BigEndian.Uint32(buf))
		entries = append(entries, entry)
	}
	return
}

func readDictContent(dict *Reader, idx *DictEntry, r io.Reader) error {
	var contents []*EntryContent
	b := make([]byte, idx.Size)
	n, err := io.ReadFull(r, b)
	if err != nil && err != io.EOF {
		return errors.Wrap(err, "failed to read entry data")
	} else if n != idx.Size {
		return errors.Errorf("entry data size not match: %v != %v", n, idx.Size)
	}
	buf := bytes.NewBuffer(b)

	for _, v := range dict.Info.SameTypeSequence {
		typ := ContentType(v)
		reader := EntryContentReaders[typ]
		if reader == nil {
			return errors.New("no reader found for type '" + string(typ) + "'")
		}

		content, err := reader(dict, idx, typ, buf)
		if err != nil {
			return errors.Wrap(err, "failed to read entry")
		}
		contents = append(contents, content)
	}
	idx.Contents = contents
	return nil
}

func (self *Reader) LoadSynonym(reader io.Reader) (err error) {
	r := bufio.NewReader(reader)
	var b []byte
	buf := make([]byte, 4)
	for {
		b, err = r.ReadBytes(0)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
		if _, err = io.ReadFull(r, buf); err != nil {
			return
		}

		i := int(binary.BigEndian.Uint32(buf))
		entry := self.Entries[i]

		if len(b) > 0 && b[len(b)-1] == 0 {
			entry.Synonyms = append(entry.Synonyms, string(b[:len(b)-1]))
		} else {
			entry.Synonyms = append(entry.Synonyms, string(b))
		}
	}
	return
}
