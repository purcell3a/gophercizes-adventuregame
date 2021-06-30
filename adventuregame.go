package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var defaultHandlerTmpl = `json:"<!DOCTYPE html>

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
                <a href="/{{.Chapter}}"></a>
                {{.Text}}
            </li>
        </ul>
    </body>

	</html>"`

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

func main() {
	// filename := flag.String("file", "story.json", "the json file")
	// flag.Parse()
	// fmt.Printf("Using the story in %s. \n", *filename)

	// f, err := os.Open(*filename)
	// if err != nil {
	// 	panic(err)
	// }
	// d := json.NewDecoder((f))
	// var story cyoa.Story
	// if err := d.Decode(&story); err != nil{
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", story)
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

	fmt.Println(story)
	
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}
