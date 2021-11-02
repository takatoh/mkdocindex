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
      {{.Style|styletag}}
  </head>
  <body>
    <header>
      <h1>{{.Title}}</h1>
    </header>
    <main>
    <div id="sidebar">
      <h2>Table of contents</h2>
      {{.TOC|safehtml}}
    </div>
    <div id="container">
      <div id="search-box">
        <input type="text" id="search-word" placeholder="ページ内テキスト検索">
        <input type="button" id="search-button" value="検索">
        <input type="button" id="reset-button" value="リセット">
      </div>
      {{.Main|safehtml}}
    </div>
    </main>
  </body>
</html>
`

	styleMonolithic = `:root {
  --header-height: 75px;
  --sidebar-width: 25%;
}

* {
  margin: 0;
}

header {
  position: sticky;
  width: 100%;
  height: var(--header-height);
  border-bottom: 1px solid grey;
}

header h1 {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  left: 5%;
}

main {
  display: flex;
  min-height: calc(100vh - var(--header-height) - 1px);
  width: 100%;
  height: 100%;
}

main div#sidebar {
  position: sticky;
  width: var(--sidebar-width);
  padding-top: 1.5em;
  padding-left: 1.5em;
  padding-right: 0.5em;
  border-right: 1px solid grey;
}

main div#sidebar h2 {
  margin-top: 0.2em;
}

main div#sidebar ul {
  padding-left: 1.5em;
  list-style: disc;
}

main div#sidebar li a {
  color: black;
  text-decoration: none;
  word-break: break-all;
}

main div#container {
  position: sticky;
  padding: 1.5em;
}

.highlight {
  font-weight: bold;
  background-color: yellow;
}
`
)

type infoMonolithic struct {
	Title string
	TOC   string
	Main  string
	Style string
}

func newInfoMonolithic(info *indexinfo.IndexInfoMonolithic) *infoMonolithic {
	p := new(infoMonolithic)
	p.Title = info.Name
	p.TOC = genTOC(info)
	p.Main = genMain(info)
	p.Style = styleMonolithic
	return p
}

func GenerateMonolithic(info *indexinfo.IndexInfoMonolithic) {
	infoM := newInfoMonolithic(info)
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
		"styletag": func(text string) template.HTML {
			return template.HTML("<style>\n" + text + "</style>")
		},
	}
	os.Remove(info.Path + "/index.html")
	t, _ := template.New("index").Funcs(funcMap).Parse(tmplMonolithic)
	w, _ := os.OpenFile(info.Path+"/index.html", os.O_WRONLY|os.O_CREATE, 0600)
	t.ExecuteTemplate(w, "index", infoM)
}

func genTOC(info *indexinfo.IndexInfoMonolithic) string {
	return "<h2>Table of contents</h2>"
}

func genMain(info *indexinfo.IndexInfoMonolithic) string {
	return "<h2>" + info.Name + "</h2>"
}
