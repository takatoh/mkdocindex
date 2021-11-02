package indexinfo

import "path/filepath"

type IndexInfo struct {
	Path        string
	Name        string
	Directories []string
	Files       []string
}

func New(path string, dirs []string, files []string) *IndexInfo {
	p := new(IndexInfo)
	p.Path = path
	p.Name = filepath.Base(path)
	for _, d := range dirs {
		p.Directories = append(p.Directories, filepath.Base(d))
	}
	for _, f := range files {
		p.Files = append(p.Files, filepath.Base(f))
	}
	return p
}

type IndexInfoMonolithic struct {
	Path        string
	Name        string
	Directories []*IndexInfoMonolithic
	Files       []string
	Level       uint8
}

func NewMonolithic(path string, name string, files []string, level uint8) *IndexInfoMonolithic {
	p := new(IndexInfoMonolithic)
	p.Path = path
	p.Name = name
	for _, f := range files {
		p.Files = append(p.Files, filepath.Base(f))
	}
	if level > 6 {
		p.Level = 0
	} else {
		p.Level = level
	}
	return p
}
