package cyoa

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func init() {
	templ = template.Must(template.New("").Parse(defaulTemp))
}

var templ *template.Template

var defaulTemp = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose Your Own Adventure</title>
</head>
<body>
	<section class="page">
    <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
    {{range .Options}}
        <li><a href="/{{.Arc}}">{{.Text}}</a></li>
    {{end}}
    </ul>
    </section>
	<style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
</body>
</html>`

func NewHandler(s Story) http.Handler {
	return handler{s}

}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if _, ok := h.s[path]; ok {
		err := templ.Execute(w, h.s[path])
		if err != nil {
			panic(err)
		}
	}
}

type Story map[string]Arc

type Arc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func ParseStory(filename string) (Story, error) {
	var story Story
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	parser := json.NewDecoder(f)
	jsonerr := parser.Decode(&story)
	if jsonerr != nil {
		return nil, jsonerr
	}
	return story, nil

}
