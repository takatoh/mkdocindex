package htmlgenerator

import (
	"html/template"
	"os"
	"strconv"

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
    {{.Script|scripttag}}
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
	scriptMonolithic = `const highlight = (elem, pat) => {
  const innerHighlight = (node, pat) => {
    let skip = 0;
    if (node.nodeType === 3) {
      const pos = node.data.toUpperCase().indexOf(pat);
      if (pos >= 0) {
        const spannode = document.createElement("span");
        spannode.className = "highlight";
        const middlebit = node.splitText(pos);
        const endbit = middlebit.splitText(pat.length);
        const middleclone = middlebit.cloneNode(true);
        spannode.appendChild(middleclone);
        middlebit.replaceWith(spannode);
        skip += 1;
      }
    } else if (node.nodeType === 1 && node.childNodes && !/script|style/i.test(node.tagName)) {
      // note: nodeType === 1 means ELEMENT_NODE.
      for (let i = 0; i < node.childNodes.length; ++i) {
        i += innerHighlight(node.childNodes[i], pat);
      }
    }
    return skip;
  };
  elem.childNodes.forEach((e) => {
    innerHighlight(e, pat.toUpperCase());
  });
};

const removeHighlight = (elem) => {
  elem.querySelectorAll("span.highlight").forEach((e) => {
    const inner = e.innerHTML;
    e.replaceWith(inner);
  });
};

const doHighlight = () => {
  const tree = document.getElementById("container");
  // remove highlights
  removeHighlight(tree);
  // split to words by whitespaces
  const words = document.getElementById("search-word").value.trim().split(/\s+/);
  // highlight words
  for (let word of words) {
    highlight(tree, word);
  }
  // scroll to first highlight
  const firstHighlight = document.getElementsByClassName("highlight")[0];
  const ypos = firstHighlight.getBoundingClientRect().top + window.pageYOffset;
  window.scrollTo({
    top: ypos,
    left: 0,
    behavior: "smooth"
  });
}

window.addEventListener("DOMContentLoaded", () => {
  document.getElementById("search-button").addEventListener("click", doHighlight, false);
  document.getElementById("reset-button").addEventListener("click", () => {
    location.reload(true);
  });
}, false);
`
)

type infoMonolithic struct {
	Title  string
	TOC    string
	Main   string
	Style  string
	Script string
}

func newInfoMonolithic(info *indexinfo.IndexInfoMonolithic) *infoMonolithic {
	p := new(infoMonolithic)
	p.Title = info.Name
	p.TOC = genTOC(info, "sec-0")
	p.Main = genMain(info, "sec-0")
	p.Style = styleMonolithic
	p.Script = scriptMonolithic
	return p
}

func GenerateMonolithic(info *indexinfo.IndexInfoMonolithic) {
	infoM := newInfoMonolithic(info)
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
		"styletag": func(text string) template.HTML {
			return template.HTML("<style>\n" + text + "</style>")
		},
		"scripttag": func(text string) template.HTML {
			return template.HTML("<script>\n" + text + "</script>")
		},
	}
	os.Remove(info.Path + "/index.html")
	t, _ := template.New("index").Funcs(funcMap).Parse(tmplMonolithic)
	w, _ := os.OpenFile(info.Path+"/index.html", os.O_WRONLY|os.O_CREATE, 0600)
	t.ExecuteTemplate(w, "index", infoM)
}

func genTOC(info *indexinfo.IndexInfoMonolithic, id string) string {
	var toc string
	anchorQuoted := "\"#" + id + "\""
	toc += "<ul>\n"
	toc += "<li><a href=" + anchorQuoted + ">" + info.Name + "</a>\n"
	if len(info.Directories) > 0 {
		for idx, dir := range info.Directories {
			toc += genTOC(dir, buildId(id, idx))
		}
	}
	toc += "</li>\n"
	toc += "</ul>\n"
	return toc
}

func genMain(info *indexinfo.IndexInfoMonolithic, id string) string {
	var content string
	if info.Level < 6 {
		h := "h" + strconv.Itoa(int(info.Level))
		attrId := "id=\"" + id + "\""
		content += "<" + h + " " + attrId + ">" + info.Name + "</" + h + ">\n"
		content += "<ul>\n"
		content += genFileList(info.Files, info.Path)
		content += "</ul>\n"
		for idx, dir := range info.Directories {
			content += genMain(dir, buildId(id, idx))
		}
	} else if info.Level == 6 {
		attrId := "id=\"" + id + "\""
		content += "<h6 " + attrId + ">" + info.Name + "</h6>\n"
		content += "<ul>\n"
		content += genFileList(info.Files, info.Path)
		content += genDirList(info.Directories, id)
		content += "</li>\n"
	} else {
		content += "<ul>\n"
		content += genDirList(info.Directories, id)
		content += "</ul>\n"
	}
	return content
}

func genFileList(files []string, parentPath string) string {
	var filelist string
	for _, filename := range files {
		anchorQuoted := "\"" + parentPath + "/" + filename + "\""
		filelist += "<li><a href=" + anchorQuoted + ">" + filename + "</a></li>\n"
	}
	return filelist
}

func genDirList(dirs []*indexinfo.IndexInfoMonolithic, idPrefix string) string {
	var dirlist string
	for idx, dir := range dirs {
		attrId := buildAttrId(idPrefix, idx)
		dirlist += "<li " + attrId + ">" + dir.Name + "\n"
		if len(dir.Files) > 0 {
			dirlist += genFileList(dir.Files, dir.Path)
		}
		if len(dir.Directories) > 0 {
			dirlist += "<ul>\n"
			dirlist += genDirList(dir.Directories, buildId(idPrefix, idx))
			dirlist += "</ul>\n"
		}
		dirlist += "</li>\n"
	}
	return dirlist
}

func buildAttrId(prefix string, idx int) string {
	return "id=\"" + buildId(prefix, idx) + "\""
}

func buildId(prefix string, idx int) string {
	return prefix + "-" + strconv.Itoa(int(idx))
}
