/*package main

import "fmt"

type Stu struct {
	name string
	age  int
}

type N int

func main() {
	var s Stu
	fmt.Println(s.setAge(11))

	var a N = 23
	a.toString()

}

func (s *Stu) setAge(a int) int {
	s.age = a + 10
	return s.age
}

func (n N) toString() {
	fmt.Printf("%s", n)
}
*/

package main

import (
	"fmt"
	"sync"
)

type N int

func (s Stu) value() {
	s.age = 11

}

func (s *Stu) pointer() {
	s.age = 11

}

type Stu struct {
	people
	age int
}

type people struct {
	name string
}

func (p *people) getname() string {
	return p.name
}

type data struct {
	sync.Mutex
	buf [1024]byte
}

func main() {
	var s1 Stu
	s1.age = 22
	s1.value()
	fmt.Println(s1.age)
	var s2 *Stu
	s2 = &Stu{age: 22, people: people{name: "gakki"}}
	s2.pointer()
	fmt.Println(s2.age)
	fmt.Println(s2.getname())

	d := data{}
	d.Lock()
	//业务处理
	d.Unlock()

	var s []int
	s = append(s, 1)
	fmt.Println(s)

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}

	x := 1
	{
		x := 2
		fmt.Print(x)
	}
	fmt.Println(x)
}
