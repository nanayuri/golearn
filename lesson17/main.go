package main

import "fmt"

type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type data struct {
}

func (d *data) test() {

}

func (d *data) string() string {
	return "aaa"
}

func pp(a stringer) {
	fmt.Println(a.string())
}

func main() {
	/*t := 3
	var f float32 = float32(t)
	fmt.Println(f)*/
	/*var t int
	t = 4

	var s string
	s = "aaa"*/

	var i interface{}
	i = 4
	fmt.Println(i)
	i = "abc"
	fmt.Println(i)

	m := make(map[string]interface{}, 0)
	m["key1"] = 0
	m["key2"] = "value2"
	s := &stu{
		name: "gakki",
		age:  11,
	}
	m["key3"] = *s
	fmt.Println(m)

	d := m["key1"].(int) //类型断言
	fmt.Println(d)
	var d1 int
	d1 = d
	fmt.Println(d1)

	var s1 stu
	s1 = m["key3"].(stu)
	fmt.Println(s1)

	aaa := m["key3"].(stu)
	aaa.age = 4
	aaa.name = "yuki"
	fmt.Println(aaa)

	d2, ok := m["key4"].(int)
	fmt.Println(d2, ok)

}

type stu struct {
	name string
	age  int
}
