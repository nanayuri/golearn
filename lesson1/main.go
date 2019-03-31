package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a,s = 1000, "sad"
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(s))

	r := 1
	fmt.Println(r)
	x,y := 1,2
	x,y = y,x
	println(x, y)
}
