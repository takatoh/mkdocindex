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
	directories []*IndexMaker
	files       []string
}

func New(path string) *IndexMaker {
	p := new(IndexMaker)
	p.path = path
	p.read()
	return p
}

func (m *IndexMaker) Make() {
	var dirs []string

	for _, d := range m.directories {
		dirs = append(dirs, d.path)
	}
	info := indexinfo.New(m.path, dirs, m.files)
	htmlgenerator.Generate(info)

	for _, d := range m.directories {
		d.Make()
	}
}

func (m *IndexMaker) read() {
	var entries []string

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
			m.directories = append(m.directories, New(f))
		} else {
			m.files = append(m.files, f)
		}
	}
}

func (m *IndexMaker) MakeMonolithic() {
	infoTree := m.ConvertInfoMonolithic(2)
	htmlgenerator.GenerateMonolithic(infoTree)
}

func (m *IndexMaker) ConvertInfoMonolithic(level uint8) *indexinfo.IndexInfoMonolithic {
	info := indexinfo.NewMonolithic(m.path, m.files, level)
	for _, d := range m.directories {
		info.Directories = append(info.Directories, d.ConvertInfoMonolithic(level+1))
	}
	return info
}
