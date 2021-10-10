package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/runes"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
)

func main() {
	sampleNopTransform()
	sampleCharCode()
	sampleWidth()
	sampleWideNarrow()
	sampleToWiden()
}

func sampleToWiden() {
	// カタカナであれば全角にする
	t := runes.If(runes.In(unicode.Katakana), width.Widen, nil)
	// ５アアAα
	fmt.Println(t.String("５ｱアAα"))
}

func sampleWideNarrow() {
	str := "５ｱアAα"
	fmt.Println("Original: " + str)
	fmt.Println("Fold ***")
	for _, r := range width.Fold.String(str) {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}
	fmt.Println("Narrow ***")
	for _, r := range width.Narrow.String(str) {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}
	fmt.Println("Widen ***")
	for _, r := range width.Widen.String(str) {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}
}

func foldShiftJISFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// Shift_JISからUTF-8にしてから全角英数などは半角に、半角カナなどは全角にする
	dec := japanese.ShiftJIS.NewDecoder()
	t := transform.Chain(dec, width.Fold)
	r := transform.NewReader(f, t)
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}

func sampleWidth() {
	rs := []rune{'5', 'ｱ', 'ア', 'A', 'α'}
	fmt.Println("rune\tWide\tNarrow\tFolded\tKind")
	fmt.Println("----------------------------------------------------")
	for _, r := range rs {
		p := width.LookupRune(r)
		w, n, f, k := p.Wide(), p.Narrow(), p.Folded(), p.Kind()
		fmt.Printf("%2c\t%2c\t%3c\t%3c\t%s\n", r, w, n, f, k)
	}
}

func sampleCharCode() {
	if err := printCSV("test.csv"); err != nil {
		log.Println(err)
	}
}

func printCSV(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	dec := japanese.ShiftJIS.NewDecoder()
	cr := csv.NewReader(dec.Reader(f))
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(rec)
	}
	return nil
}

func sampleNopTransform() {
	r := strings.NewReader("Hello, World")
	tw := transform.NewWriter(os.Stdout, transform.Nop)
	if _, err := io.Copy(tw, r); err != nil {
		log.Fatalln(err)
	}
	log.Println()
	t := transform.Chain(transform.Nop, transform.Discard)
	tw = transform.NewWriter(os.Stdout, t)
	if _, err := io.Copy(tw, r); err != nil {
		log.Fatalln(err)
	}
	log.Println()
}
