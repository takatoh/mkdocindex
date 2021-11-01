package htmlgenerator

import (
	"html/template"
	"os"

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

func Generate(info *indexinfo.IndexInfo) {
	os.Remove(info.Path + "/index.html")
	t, _ := template.New("index").Parse(tmpl)
	w, _ := os.OpenFile(info.Path+"/index.html", os.O_WRONLY|os.O_CREATE, 0600)
	t.ExecuteTemplate(w, "index", info)
}

const (
	tmplMonolithic = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>{{.Title}}</title>
  </head>
  <body>
    <header>
      <h1>{{.Title}}</h1>
    </header>
    <main>
    <div id="sidebar">
      <h2>Table of contents</h2>
      {{.TOC}}
    </div>
    <div id="container">
      <div id="search-box">
        <input type="text" id="search-word" placeholder="ページ内テキスト検索">
        <input type="button" id="search-button" value="検索">
        <input type="button" id="reset-button" value="リセット">
      </div>
      {{.Main}}
    </div>
    </main>
  </body>
</html>
`
)

type infoMonolithic struct {
	Title string
	TOC   string
	Main  string
}

func newInfoMonolithic(info *indexinfo.IndexInfoMonolithic) *infoMonolithic {
	p := new(infoMonolithic)
	p.Title = info.Name
	p.TOC = genTOC(info)
	p.Main = genMain(info)
	return p
}

func GenerateMonolithic(info *indexinfo.IndexInfoMonolithic) {
	infoM := newInfoMonolithic(info)
	os.Remove(info.Path + "/index.html")
	t, _ := template.New("index").Parse(tmplMonolithic)
	w, _ := os.OpenFile(info.Path+"/index.html", os.O_WRONLY|os.O_CREATE, 0600)
	t.ExecuteTemplate(w, "index", infoM)
}

func genTOC(info *indexinfo.IndexInfoMonolithic) string {
	return "<h2>Table of contents</h2>"
}

func genMain(info *indexinfo.IndexInfoMonolithic) string {
	return "<h2>" + info.Name + "</h2>"
}
