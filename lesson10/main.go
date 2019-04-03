package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "ab c\x61\u0041"
	fmt.Println(s)
	var s2 string
	fmt.Println(s2 == "")
	s3 := `abc\ndef`
	//原生字符串`
	fmt.Println(s3)
	s4 := "ab" + "cd"
	fmt.Println(s4 == "abcd")
	fmt.Println(s4 > "abc")
	s5 := "abcde"
	//fmt.Printf("%c", s5[1])
	s6 := s5[1:3]
	fmt.Println(s6)
	s7 := s5[:4]
	fmt.Println(s7)

	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer((&s5))))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer((&s6))))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer((&s6))))

	for i := 0; i < len(s5); i++ {
		fmt.Printf("%c", s5[i])
	}

	for i, e := range s5 {
		fmt.Printf("%d, %c\n", i, e)
	}

	b := []byte(s5)
	fmt.Println(b)
	b[1] = 99
	s8 := string(b)
	fmt.Println(s8)

	//s9:=make([]byte,100)
}
