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
	entries, _ := filepath.Glob(i.path + "/*")

	for _, f := range entries {
		finfo, _ := os.Stat(f)
		if finfo.IsDir() {
			i.directories = append(i.directories, f)
		} else {
			i.files = append(i.files, f)
		}
	}

	fmt.Println("Directories:")
	for _, d := range i.directories {
		fmt.Println(d)
	}
	fmt.Println("")
	fmt.Println("Files:")
	for _, f := range i.files {
		fmt.Println(f)
	}
}
