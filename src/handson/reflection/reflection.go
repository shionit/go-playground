package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/google/go-cmp/cmp"

	"golang.org/x/xerrors"
)

func main() {
	sampleValue()
	sampleSetValue()
	sampleDeepEqual()
	sampleType()
}

func sampleType() {
	t1 := reflect.TypeOf(100)
	fmt.Println(t1) // int
	t2 := reflect.ValueOf("hello").Type()
	fmt.Println(t2) // string

	v1 := reflect.New(reflect.TypeOf(0))
	fmt.Println(v1.Kind(), v1.Elem().Kind(), v1.Elem())

	var n int
	v2 := reflect.NewAt(reflect.TypeOf(0), unsafe.Pointer(&n))
	v2.Elem().SetInt(200)
	fmt.Println(n) // 200
}

func sampleDeepEqual() {
	type A struct {
		N int
	}
	type B struct {
		A *A
		M int
	}
	b1 := &B{A: &A{N: 100}, M: 200}
	b2 := &B{A: &A{N: 100}, M: 200}

	fmt.Println("==", b1 == b2)                         // false
	fmt.Println("DeepEqual", reflect.DeepEqual(b1, b2)) // true
	fmt.Println("cmp.Equal", cmp.Equal(b1, b2))         // true
}

func sampleSetValue() {
	var s string
	err := setString(&s, "hoge")
	fmt.Println("result: ", s, err == nil) // hoge true

	err = setString(s, "fuga")
	fmt.Println("result: ", s, err == nil) // hoge false

	var n int
	err = setString(&n, "hoge")
	fmt.Println("result: ", n, err == nil) // 0 false
}

func setString(val interface{}, str string) error {
	v := reflect.ValueOf(val)
	fmt.Println(v.Kind())
	if v.Kind() != reflect.Ptr {
		return xerrors.New("val is not string pointer")
	}
	v = v.Elem()
	fmt.Println(v.Kind())
	if v.Kind() != reflect.String {
		return xerrors.New("val is not string pointer")
	}
	if v.CanSet() {
		v.Set(reflect.ValueOf(str))
	}
	return nil
}

func sampleValue() {
	var data interface{} = 100
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Int:
		fmt.Println(v.Kind(), v.Int())
	case reflect.String:
		fmt.Println(v.Kind(), v.String())
	case reflect.Bool:
		fmt.Println(v.Kind(), v.Bool())
	}
}
