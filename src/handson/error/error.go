package main

import (
	"fmt"
	"os"
)

type ToStringerError string

func (e ToStringerError) Error() string {
	return string(e)
}

type Stringer interface {
	String() string
}

func ToStringer(v interface{}) (Stringer, error) {
	str, ok := v.(Stringer)
	if !ok {
		return nil, ToStringerError("v is not Stringer")
	}
	return str, nil
}

type A int
type B string

func (b B) String() string {
	return "BBB"
}

func main() {
	if s, err := ToStringer(A(3)); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		println(s.String())
	}

	if s, err := ToStringer(B("string")); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		println(s.String())
	}
}
