package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/goccy/go-json"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("src/handson/json-read/input.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer func() {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("Error closing JSON file:", err)
		}
	}()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	var post Post
	jsonErr := json.Unmarshal(jsonData, &post)
	if jsonErr != nil {
		fmt.Println("Error unmarshalling JSON:", jsonErr)
		return
	}
	fmt.Println(post)
}
