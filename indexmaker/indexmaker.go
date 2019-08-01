package indexmaker

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)


const (
	tmpl = `<!DOCTYPE html>
<html>
  <head>
    <title>{{.Name}}</title>
  </head>
  <body>
    <h1>{{.Name}}</h1>

    {{if gt (len .Directories) 0}}
    <h2>Directories</h2>
    <ul>
      {{range $i, $v := .Directories}}
      <li><a href="{{index $.Directories $i}}/index.html">{{index $.Directories $i}}</a></li>
      {{end}}
    </ul>
    {{end}}

    {{if gt (len .Files) 0}}
    <h2>Files</h2>
    <ul>
      {{range $i, $v := .Files}}
      <li><a href="{{index $.Files $i}}">{{index $.Files $i}}</a></li>
      {{end}}
    </ul>
    {{end}}
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

	for _, d := range m.Directories {
		maker := New(d)
		maker.Make()
	}
}

func (m *IndexMaker) getEntries() {
	var entries []string
	var directories []string
	var files []string

	ents, _ := filepath.Glob(m.Path + "/*")
	for _, e := range ents {
		e2 := filepath.Base(e)
		if strings.Index(e2, ".") != 0 && e2 != "index.html" && strings.Index(e2, ".exe") < 0 {
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
	os.Remove(m.Path + "/index.html")
	t, _ := template.New("index").Parse(tmpl)
	w, _ := os.OpenFile(m.Path + "/index.html", os.O_WRONLY|os.O_CREATE, 0600)
	t.ExecuteTemplate(w, "index", newIndexInfo(m))
}


type IndexInfo struct {
	Name        string
	Directories []string
	Files       []string
}

func newIndexInfo(m *IndexMaker) *IndexInfo {
	p := new(IndexInfo)
	p.Name = filepath.Base(m.Path)
	for _, d := range m.Directories {
		p.Directories = append(p.Directories, filepath.Base(d))
	}
	for _, f := range m.Files {
		p.Files = append(p.Files, filepath.Base(f))
	}
	return p
}
