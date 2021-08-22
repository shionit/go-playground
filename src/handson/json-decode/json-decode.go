package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	json2 "github.com/goccy/go-json"
)

const (
	TryCount = 10000
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := &Person{Name: "GoGo", Age: 40}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	decodeJson(buf.String())
	decodeJson2(buf.String())
}

func decodeJson(str string) {
	var p2 Person
	timeStart := time.Now()
	for i := 0; i < TryCount; i++ {
		buf := bytes.NewBufferString(str)
		dec := json.NewDecoder(buf)
		if err := dec.Decode(&p2); err != nil {
			log.Fatal(err)
		}
		fmt.Println(p2)
	}
	timeEnd := time.Now()
	fmt.Print("encoding/json ")
	fmt.Printf("Execute time: %.3f [s]\n\n", timeEnd.Sub(timeStart).Seconds())
}

func decodeJson2(str string) {
	var p2 Person
	timeStart := time.Now()
	for i := 0; i < TryCount; i++ {
		buf := bytes.NewBufferString(str)
		dec := json2.NewDecoder(buf)
		if err := dec.Decode(&p2); err != nil {
			log.Fatal(err)
		}
		fmt.Println(p2)
	}
	timeEnd := time.Now()
	fmt.Print("goccy/go-json ")
	fmt.Printf("Execute time: %.3f [s]\n\n", timeEnd.Sub(timeStart).Seconds())
}
