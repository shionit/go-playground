package main

import (
	"fmt"
	"io"
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
	jsonFile, err := os.Open("src/handson/json-decoder/input.json")
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

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}
}
