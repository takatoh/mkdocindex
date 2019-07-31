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

func (i *IndexMaker) Make() {
	directories, files := i.getEntries()

	i.makeIndex(directories, files)
}

func (m *IndexMaker) getEntries() ([]string, []string) {
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

	return directories, files
}

func (m *IndexMaker) makeIndex(directories, files []string) {
	fmt.Println(m.path)
	fmt.Println("")
	fmt.Println("Directories:")
	for _, d := range directories {
		fmt.Println(d)
	}
	fmt.Println("")
	fmt.Println("Files:")
	for _, f := range files {
		fmt.Println(f)
	}
}
