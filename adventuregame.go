package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var chapters Story

type Story struct {
	Intro []string
	NewYork []string
	Debate []string
}

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
