package indexmaker

import (
	"os"
	"path/filepath"
	"strings"
	"html/template"
)


const (
	tmpl = `<!DOCTYPE html>
<html>
  <head>
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
	m.makeIndex()

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

func (m *IndexMaker) makeIndex() {
	os.Remove(m.path + "/index.html")
	t, _ := template.New("index").Parse(tmpl)
	w, _ := os.OpenFile(m.path + "/index.html", os.O_WRONLY|os.O_CREATE, 0600)
	t.ExecuteTemplate(w, "index", newIndexInfo(m))
}


type IndexInfo struct {
	Name        string
	Directories []string
	Files       []string
}

func newIndexInfo(m *IndexMaker) *IndexInfo {
	p := new(IndexInfo)
	p.Name = filepath.Base(m.path)
	for _, d := range m.directories {
		p.Directories = append(p.Directories, filepath.Base(d))
	}
	for _, f := range m.files {
		p.Files = append(p.Files, filepath.Base(f))
	}
	return p
}
