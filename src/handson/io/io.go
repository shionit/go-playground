package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	copySample()
	seekerSample()
	pipeSample()
	limitSample()
	multiWriteSample()
	multiReadSample()
	teeReaderSample()
}

func teeReaderSample() {
	var buf bytes.Buffer
	r := strings.NewReader("Hello, tee世界\n")
	tee := io.TeeReader(r, &buf)
	if _, err := io.Copy(os.Stdout, tee); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func multiReadSample() {
	r1 := strings.NewReader("Hello, ")
	r2 := strings.NewReader("世界\n")
	r := io.MultiReader(r1, r2)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func multiWriteSample() {
	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2)
	if _, err := fmt.Fprint(w, "Hello, 世界Multi"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("buf1: " + buf1.String())
	fmt.Println("buf2: " + buf2.String())
}

func limitSample() {
	r := io.LimitReader(
		strings.NewReader("Hello, 世界"), 5,
	)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func pipeSample() {
	r, w := io.Pipe()
	go func() {
		if _, err := fmt.Fprintln(w, "Hello, 世界!"); err != nil {
			log.Fatal(err)
		}
		if err := w.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func seekerSample() {
	r := strings.NewReader("Hello, 世界")
	if _, err := r.Seek(2, io.SeekStart); err != nil {
		log.Fatal(err)
	}
	if _, err := io.CopyN(os.Stdout, r, 2); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	if _, err := r.Seek(-4, io.SeekCurrent); err != nil {
		log.Fatal(err)
	}
	if _, err := io.CopyN(os.Stdout, r, 7); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	if _, err := r.Seek(-6, io.SeekEnd); err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func copySample() {
	r1 := strings.NewReader("Hello, 世界")
	if _, err := io.Copy(os.Stdout, r1); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	r2 := strings.NewReader("Hello, 世界")
	if _, err := io.CopyN(os.Stdout, r2, 5); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
