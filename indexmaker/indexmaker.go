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
	w, _ := os.OpenFile("index.html", os.O_WRONLY|os.O_CREATE, 0600)

	fmt.Fprintln(w, m.path)
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Directories:")
	for _, d := range m.directories {
		fmt.Fprintln(w, d)
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Files:")
	for _, f := range m.files {
		fmt.Fprintln(w, f)
	}
}
