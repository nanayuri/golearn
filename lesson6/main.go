package main

import (
	"errors"
	"fmt"
)

func check(x int) error {
	if x < 0 {
		return errors.New("x<0")
	}
	return nil
}

func main() {
	//x:=3
	/*if x := 8; x > 5 {
		fmt.Println(x)
	} else if x > 7 {
		fmt.Println(x)
	}*/
	x := -1
	if err := check(x); err == nil {
		x++
		fmt.Println(x)
	} else {
		fmt.Println(err)
	}
	for i := 0; i < 4; i++ {
		if i > 2 {
			break
		}
		fmt.Println(i)
	}
	//map channel slice array:range
	var data [3]int = [3]int{10, 20, 30}
	for i := range data {
		fmt.Println(i, data[i])
	}
	data2 := [3]string{"aa", "bb", "cc"}
	var data2Copy [3]*string
	for i, s := range data2 {
		//fmt.Println(&i, &s)
		data2Copy[i] = &s
	}
	for _, s := range data2Copy {
		fmt.Println(*s)
	}

	//goto break continue

	for i := 0; i < 4; i++ {
		fmt.Println(i)
		if i > 2 {
			goto end
		}
	end:
		fmt.Println("end")
	}

	testSwitch()

}
func testSwitch() {
	a, b, c, x := 1, 2, 3, 2
	switch x {
	default:
		fmt.Println("all not equal")
	case a:
		fmt.Println("x=a")
	case b:
		fmt.Println("x=b")
	case c:
		fmt.Println("x=c")

	}
}
