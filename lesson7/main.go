package main

import "fmt"

func main() {
	a := sayHello
	b := sayHello2
	a = b
	a()
	if a == nil {
		fmt.Println("nil")
	}

	var a1 *int = sayHello3()
	fmt.Println(a1, *a1)
	pkg.Smile()

}

//具有相同参数和返回值的函数才是同一类型
func sayHello() {
	fmt.Println("hello")
	a := func() { fmt.Println("222") }
	a()
}

func sayHello2() {
	fmt.Println("hello")
	a := func() { fmt.Println("222") }
	a()
}

func sayHello3() *int {
	a := 100
	return &a
}
