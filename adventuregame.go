package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type handler struct {
	s Story
	t *template.Template
}

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template

var defaultHandlerTmpl = `<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <title>Adventure Game</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
        <p>{{.}}</p>
        {{end}}
        <ul>
            {{range .Options}}
            <li>
				<a href="/{{.Chapter}}">
				{{.Text}}
				</a>
            </li>
            {{end}}
        </ul>
    </body>
</html>`

func NewHandler(s Story, tmpl *template.Template) http.Handler {
	if tmpl == nil {
		tmpl = tpl
	}
	return handler{s, tpl}
}

// func NewHandler(s Story) http.Handler {
// 	return handler{s}
// }

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	//defualt to start of story if no path
	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "chapter not found.", http.StatusNotFound)
}

func main() {
	port := flag.Int("port", 3000, "the port to start the server")

	jsonFile, err := os.Open("story.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("running", jsonFile)

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	var story Story

	jsonErr := json.Unmarshal(data, &story)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	h := NewHandler(story, tpl)
	fmt.Printf("starting the on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
