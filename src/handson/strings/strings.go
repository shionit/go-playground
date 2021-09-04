package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	strconvSample()
	stringsSample()
	replaceSample()
}

func replaceSample() {
	fmt.Println(strings.Replace("郷に入っては郷に従え", "郷", "Go", 1))
	fmt.Println(strings.ReplaceAll("郷に入っては郷に従え", "郷", "Go"))

	r := strings.NewReplacer(
		"郷", "Go",
		"従え", "従おう",
	)
	fmt.Println(r.Replace("郷に入っては郷に従え"))

	toUpper := func(r rune) rune {
		if 'a' <= r && r <= 'z' {
			return r - ('a' - 'A')
		}
		return r
	}
	fmt.Println(strings.Map(toUpper, "Hello, world."))

	fmt.Println(strings.ToUpper("Hello, world."))
	fmt.Println(strings.ToLower("Hello, world."))

	b := []byte{0x0A, 0x0B, 0x0C}
	fmt.Println(bytes.ReplaceAll(b, []byte{0x0B}, []byte{0xFF}))
}

func stringsSample() {
	fmt.Println(strings.Split("a b c", " "))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, ","))
	fmt.Println(strings.Repeat("hoge", 3))
	fmt.Println(strings.HasPrefix("hoge_huga", "hoge"))
}

func strconvSample() {
	fmt.Println(strconv.Atoi("100"))
	fmt.Println(strconv.Itoa(100))
	fmt.Println(strconv.FormatInt(100, 16))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseInt("100", 10, 64))
}
