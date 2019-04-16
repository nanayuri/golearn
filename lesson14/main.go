package main

import (
	"fmt"
	"unsafe"
)

type node struct {
	id int32
	//next *node
	//_    int
	age int8
	tel int16
}

type user struct {
	name string
	age  int
}
type file struct {
	name string
	attr
}
type attr struct {
	owner int
	perm  int
}

func main() {
	//fmt.Println(unsafe.Sizeof(node))
	u := user{name: "gakki", age: 12}
	fmt.Println(u)
	u2 := user{"yuki", 12}
	fmt.Println(u2)

	u3 := struct {
		name string
		age  int
	}{
		name: "saki", age: 14,
	}
	fmt.Println(u3)

	/*f := file{
		name: "kuraki",
	}
	f.attr.owner = 1
	f.attr.perm = 2*/

	//fmt.Println(f)
	f2 := file{
		name: "kuraki",
		attr: attr{
			owner: 1,
			perm:  2,
		},
	}
	fmt.Println(f2)
	var a struct{}
	var b [100]struct{}
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
}
