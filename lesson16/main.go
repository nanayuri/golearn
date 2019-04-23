package main

import "fmt"

//接口是方法声明的集合
//接口是协议，任何实现接口里方法的对象，都可以称作实现了这个接口
type tester interface {
	test()
	string() string
}

type data struct {
}

func (d *data) test() {
	fmt.Println("test")
}

func (d *data) string() string {
	fmt.Println("string")
	return "abc"
}

func main() {
	d := &data{}
	//var d data
	var i tester
	i = d
	i.test()
	i.string()
}
