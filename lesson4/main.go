package main

import (
	"fmt"
	"math"
)

type Student struct {
	Age int
}

func main() {

	/*
		bool()
		byte()
		int 8 16 32 64
		uint
		float 8 16 32 64
		complex64 128 1+2i
		uintptr 指针
		string
		array
		struct
		function
		interface
		map
		slice
		channel
	*/
	s := "0Abc"
	fmt.Println([]byte(s))

	stu := new(Student)
	fmt.Println(*stu)

	fmt.Println(math.MinInt8, math.MaxInt16)

	m := make(map[int]string)
	m[1] = "1"
	m[2] = "2"
	fmt.Println(m)

	s1 := make([]int, 5)
	s1 = append(s1, 12)
	fmt.Println(s1)

	c := make(chan int, 2)
	c <- 1
	fmt.Println(<-c)

	/*p:=new(map[int]string)
	m1:=*p
	m1[1]="1"
	fmt.Println(m1)*/

	a := 10
	b := byte(a)
	fmt.Println(b)

	type color int
	var red color
	red = 1
	var green color
	green = 2
	fmt.Println(red, green)

	red1 := 1
	red1 = int(red)
	fmt.Println(red1)

}
