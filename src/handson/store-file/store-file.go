package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// WriteFile / ReadFile
	data := []byte("Hello, world\n")
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	// Create / Write / Open / Read
	file1, _ := os.Create("data2")
	defer func() {
		if err := file1.Close(); err != nil {
			panic(err)
		}
	}()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer func() {
		if err := file2.Close(); err != nil {
			panic(err)
		}
	}()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
