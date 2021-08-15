package main

import "fmt"

type Stringer interface {
	String() string
}

type StrA int

func (s StrA) String() string {
	return fmt.Sprintf("%d", s)
}

type StrB bool

func (s StrB) String() string {
	if s {
		return "true"
	}
	return "false"
}

type StrC string

func (s StrC) String() string {
	return string(s)
}

func main() {
	stringer := StrA(1)
	resolve(stringer)
	resolve(StrB(true))
	resolve(StrC("hoge"))
}

func resolve(stringer Stringer) {
	switch v := stringer.(type) {
	case StrA:
		fmt.Println("StrA", v.String())
	case StrB:
		fmt.Println("StrB", v.String())
	case StrC:
		fmt.Println("StrC", v.String())
	default:
		fmt.Println("Anyone else")
	}
}
