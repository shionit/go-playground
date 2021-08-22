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

	encodeJson(p)
	encodeJson2(p)
}

func encodeJson(p *Person) {
	timeStart := time.Now()
	for i := 0; i < TryCount; i++ {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(p); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())
	}
	timeEnd := time.Now()
	fmt.Print("encoding/json ")
	fmt.Printf("Execute time: %.3f [s]\n\n", timeEnd.Sub(timeStart).Seconds())
}

func encodeJson2(p *Person) {
	timeStart := time.Now()
	for i := 0; i < TryCount; i++ {
		var buf bytes.Buffer
		enc := json2.NewEncoder(&buf)
		if err := enc.Encode(p); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())
	}
	timeEnd := time.Now()
	fmt.Print("goccy/go-json ")
	fmt.Printf("Execute time: %.3f [s]\n\n", timeEnd.Sub(timeStart).Seconds())
}
