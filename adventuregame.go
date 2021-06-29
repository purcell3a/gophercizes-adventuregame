package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Intro struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
} `json:"intro"`


type Option struct {
	OptionName string        `json:"option name"`
	Title      string        `json:"title"`
	Story      []string      `json:"story"`
	Options    []interface{} `json:"options"`
} `json:"option"`


func main() {
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

	jsonErr := json.Unmarshal(data, &chapters)

	if jsonErr != nil{
		log.Fatal(jsonErr)
	}

	fmt.Println(chapters)
}
