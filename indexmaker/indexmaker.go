package indexmaker

import (
	"fmt"
	"os"
	"path/filepath"
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
	m.makeIndex()
}

func (m *IndexMaker) getEntries() {
	var directories []string
	var files []string

	entries, _ := filepath.Glob(m.path + "/*")

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

func (m *IndexMaker) makeIndex() {
	fmt.Println(m.path)
	fmt.Println("")
	fmt.Println("Directories:")
	for _, d := range m.directories {
		fmt.Println(d)
	}
	fmt.Println("")
	fmt.Println("Files:")
	for _, f := range m.files {
		fmt.Println(f)
	}
}
