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

func New(path) {
	p := new(IndexMaker)
	p.path = path
	return p
}

func (i *IndexMaker) Make {
	entries, _ := filepath.Glob(path + "/*")

	for _, f := range entries {
		finfo, _ := os.Stat(f)
		if finfo.IsDir() {
			directories = append(directories, f)
		} else {
			files = append(files, f)
		}
	}

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
