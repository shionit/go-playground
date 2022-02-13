package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("src/handson/xml-decoder/post.xml")
	if err != nil {
		fmt.Println("Error opening XML File:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				err := decoder.DecodeElement(&comment, &se)
				if err != nil {
					fmt.Println("Error decoding XML Element:", err)
					return
				}
				fmt.Println(comment)
			}
		}
	}
}
