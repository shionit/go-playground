package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080?p=Gopher")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { err := resp.Body.Close(); log.Fatal(err) }()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
