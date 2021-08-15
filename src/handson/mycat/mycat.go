package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var displayLineNum bool
	flag.BoolVar(&displayLineNum, "n", false, "display line number")
	flag.Parse()
	files := flag.Args()

	var lineNum uint = 0
	for _, fn := range files {
		err := readfile(fn, displayLineNum, &lineNum)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func readfile(fileName string, displayLineNum bool, u *uint) error {
	fr, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	defer fr.Close()

	scanner := bufio.NewScanner(fr)
	for scanner.Scan() {
		line := scanner.Text()
		if displayLineNum {
			*u++
			fmt.Printf("%d: %s\n", *u, line)
		} else {
			fmt.Println(line)
		}
	}
	return scanner.Err()
}
