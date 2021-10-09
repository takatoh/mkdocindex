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
