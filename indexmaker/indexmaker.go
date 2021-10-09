package indexmaker

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/takatoh/mkdocindex/htmlgenerator"
	"github.com/takatoh/mkdocindex/indexinfo"
)

const (
	tmpl = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>{{.Name}}</title>
  </head>
  <body>
    <h1>{{.Name}}</h1>

    {{if and (eq (len .Directories) 0) (eq (len .Files) 0)}}
    <p><b>No items</b></p>
    {{end}}

    {{if gt (len .Directories) 0}}
    <h2>Directories</h2>
    <ul>
      {{range $i, $v := .Directories}}
      <li><a href="{{$v}}/index.html">{{$v}}</a></li>
      {{end}}
    </ul>
    {{end}}

    {{if gt (len .Files) 0}}
    <h2>Files</h2>
    <ul>
      {{range $i, $v := .Files}}
      <li><a href="{{$v}}">{{$v}}</a></li>
      {{end}}
    </ul>
    {{end}}
  </body>
</html>
`
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
	htmlgenerator.MakeIndex(info)

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
