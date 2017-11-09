package stardictfmt

import (
	"compress/gzip"
	"github.com/wenerme/letsgo/fs"
	"github.com/wenerme/letsgo/time"
	"io"
	"os"
	"path/filepath"
	"time"
)

func LookupFile(dir string, fn string, exts ...string) (string, bool) {
	/*
		1) gStarDictDataDir + "/dic",
		2) "/usr/share/stardict/dic",
		3) g_get_home_dir() + "/.stardict/dic".
	*/
	return "", false
}

func openFile(fn string) (r io.ReadCloser, err error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			f.Close()
		}
	}()
	switch filepath.Ext(fn) {
	case ".dz":
		fallthrough
	case ".gz":
		r, err = gzip.NewReader(f)
	default:
		r = f
	}
	return
}

func findFile(base string, exts ...string) string {
	for _, v := range exts {
		path := base + v
		if wfs.Exists(path) {
			return path
		}
	}
	return ""
}

func parseDate(value string) (time.Time, error) {
	return wtime.ParseLayout(value, "2006.1.2", "2006-01-02", "Jan 02 2006")
}
