package indexmaker

import (
//	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)


const (
	tmpl = `<!DOCTYPE html>
<html>
  <head>
    <title>Index of directory</title>
  </head>
  <body>
    <h1>{{.Path}}</h1>

    <h2>Directories</h2>
    <ul>
    {{range $i, $v := .Directories}}
      <li><a href="{{index $.Directories $i}}/index.html">{{index $.Directories $i}}</a></li>
    {{end}}
    </ul>

    <h2>Files</h2>
    <ul>
    {{range $i, $v := .Files}}
      <li><a href="{{index $.Files $i}}">{{index $.Files $i}}</a></li>
    {{end}}
    </ul>
  </body>
</html>
`
)


type IndexMaker struct {
	Path        string
	Directories []string
	Files       []string
}

func New(path string) *IndexMaker {
	p := new(IndexMaker)
	p.Path = path
	return p
}

func (m *IndexMaker) Make() {
	m.getEntries()
	m.makeIndex()
}

func (m *IndexMaker) getEntries() {
	var entries []string
	var directories []string
	var files []string

	ents, _ := filepath.Glob(m.Path + "/*")
	for _, e := range ents {
		if strings.Index(e, ".") != 0 && strings.Index(e, "index.html") < 0 && strings.Index(e, ".exe") < 0 {
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

	m.Directories = directories
	m.Files = files
}

func (m *IndexMaker) makeIndex() {
	t, _ := template.New("index").Parse(tmpl)
	w, _ := os.OpenFile("index.html", os.O_WRONLY|os.O_CREATE, 0600)
	t.ExecuteTemplate(w, "index", m)
}
