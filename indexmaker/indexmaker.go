package indexmaker

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/takatoh/mkdocindex/htmlgenerator"
	"github.com/takatoh/mkdocindex/indexinfo"
)

type IndexMaker struct {
	path        string
	directories []string
	files       []string
}

func New(path string) *IndexMaker {
	p := new(IndexMaker)
	p.path = path
	return p
}

func (m *IndexMaker) Make() {
	m.getEntries()
	info := indexinfo.New(m.path, m.directories, m.files)
	htmlgenerator.Generate(info)

	for _, d := range m.directories {
		maker := New(d)
		maker.Make()
	}
}

func (m *IndexMaker) getEntries() {
	var entries []string
	var directories []string
	var files []string

	ents, _ := filepath.Glob(m.path + "/*")
	for _, e := range ents {
		e2 := filepath.Base(e)
		if strings.Index(e2, ".") != 0 && e2 != "index.html" && e2 != "mkdocindex.exe" {
			entries = append(entries, e)
		}
	}

	for _, f := range entries {
		finfo, _ := os.Stat(f)
		if finfo.IsDir() {
			directories = append(directories, f)
		} else {
			files = append(files, f)
		}
	}

	m.directories = directories
	m.files = files
}
