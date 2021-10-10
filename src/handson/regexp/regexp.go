package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+]$`)

func main() {
	sampleCompile()
	sampleMatch()
	sampleFind()
	sampleFindIndex()
	sampleFindSubmatch()
	sampleExpandString()
	sampleReplace()
	sampleReplaceLiteralString()
	sampleReplaceAllStringFunc()
}

func sampleReplaceAllStringFunc() {
	fmt.Println("sampleReplaceAllStringFunc")
	rex, err := regexp.Compile(`(^|_)[a-zA-Z]`)
	if err != nil {
		log.Fatalln(err)
	}
	s := rex.ReplaceAllStringFunc("hello_world", func(s string) string {
		return strings.ToUpper(strings.TrimLeft(s, "_"))
	})
	fmt.Println(s)
}

func sampleReplaceLiteralString() {
	fmt.Println("sampleReplaceLiteralString")
	rex, err := regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		log.Fatalln(err)
	}
	s := rex.ReplaceAllLiteralString("1986年01月12日", "${2}/${3} ${1}")
	fmt.Println(s)
}

func sampleReplace() {
	fmt.Println("sampleReplace")
	rex, err := regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		log.Fatalln(err)
	}
	s := rex.ReplaceAllString("1986年01月12日", "${2}/${3} ${1}")
	fmt.Println(s)
}

func sampleExpandString() {
	fmt.Println("sampleExpandString")
	rex, err := regexp.Compile(`(?P<Y>\d+)年(?P<M>\d+)月(?P<D>\d+)日`)
	if err != nil {
		log.Fatalln(err)
	}
	content := "1986年01月12日\n2020年03月24日"
	template := "$Y/$M/$D\n"
	var result []byte
	for _, submatches := range rex.FindAllStringSubmatchIndex(content, -1) {
		result = rex.ExpandString(result, template, content, submatches)
	}
	fmt.Printf("%q\n", result)
}

func sampleFindSubmatch() {
	fmt.Println("sampleFindSubmatch")
	rex, err := regexp.Compile(`(\d+)[^\d]`)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%q\n", rex.FindSubmatch([]byte("1986年01月12日")))
	fmt.Printf("%q\n", rex.FindStringSubmatch("1986年01月12日"))
	fmt.Printf("%q\n", rex.FindAllSubmatch([]byte("1986年01月12日"), -1))
	fmt.Printf("%q\n", rex.FindAllStringSubmatch("1986年01月12日", -1))
	fmt.Println(rex.FindSubmatchIndex([]byte("1986年01月12日")))
	fmt.Println(rex.FindStringSubmatchIndex("1986年01月12日"))
	fmt.Println(rex.FindAllSubmatchIndex([]byte("1986年01月12日"), -1))
	fmt.Println(rex.FindAllStringSubmatchIndex("1986年01月12日", -1))
}

func sampleFindIndex() {
	fmt.Println("sampleFindIndex")
	rex, err := regexp.Compile(`\d+`)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rex.FindIndex([]byte("1986年01月12日")))
	fmt.Println(rex.FindAllIndex([]byte("1986年01月12日"), -1))
	fmt.Println(rex.FindStringIndex("1986年01月12日"))
	fmt.Println(rex.FindAllStringIndex("1986年01月12日", -1))
}

func sampleFind() {
	fmt.Println("sampleFind")
	rex, err := regexp.Compile(`\d+`)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%q\n", rex.Find([]byte("1986年01月12日")))
	fmt.Printf("%q\n", rex.FindAll([]byte("1986年01月12日"), -1))
	fmt.Printf("%q\n", rex.FindString("1986年01月12日"))
	fmt.Printf("%q\n", rex.FindAllString("1986年01月12日", -1))
}

func sampleMatch() {
	fmt.Println("sampleMatch")
	rex, err := regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rex.Match([]byte("1986年01月12日")))
	fmt.Println(rex.MatchString("1986年01月12日"))
	r := strings.NewReader("1986年01月12日")
	fmt.Println(rex.MatchReader(r))
}

func sampleCompile() {
	fmt.Println("sampleCompile")
	fmt.Println(validID.MatchString("test[35]"))
	validID2, err := regexp.Compile(`^[a-z]+\[[0-9]+]$`)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(validID2.MatchString("test[35]"))
}
