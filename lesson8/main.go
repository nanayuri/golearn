package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	a := 100
	p := &a
	fmt.Println(&a)
	fmt.Printf("%p,%v\n", &p, p)
	test(p)
	var a1 = []int{1, 2, 3, 4}
	//test2("aaa",a1[:]...)
	fmt.Println(reflect.TypeOf(a1))
	b1 := a1[:]
	fmt.Println(reflect.TypeOf(b1))
	c, err := test3(1)
	fmt.Println(c, err)
}

//go的函数传值还是传指针？ 传值！！！
/*func A(a int,b bool){

}*/

func test(x *int) {
	fmt.Printf("%p, %v\n", &x, x)
}

func test2(s string, a ...int) {
	fmt.Printf("%T,%v\n", a, a)
}

func test3(a int) (int, error) {
	b := a / 3
	return b, errors.New("error")
}
