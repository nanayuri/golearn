package main

import (
	"errors"
	"fmt"
)

func test(f func(s string) (string, error)) string {
	a, _ := f("gakki")
	return a
}

func test2(s string) (string, error) {
	//fmt.Println(s)
	return s, errors.New("error")
}

func main() {
	/*func(s string) {
		fmt.Println(s)
	}("pigungun")
	//a("pigun")*/
	/*result := test(test2)
	fmt.Println(result)*/

	/*f, err := os.Open("filename")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var c []byte
	b := bytes.NewBuffer(c)
	content, _ := f.Read(b.Bytes())
	fmt.Println(content)*/

	x, y := 1, 2
	/*defer func(a int) {
		fmt.Println(a, x, y)
	}(11)*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

		}
	}()
	fmt.Println(x, y+1)

	fmt.Println(testDefer())

	//panic recover
	panic("database is not connected")
	fmt.Println("test panic")
}

func testDefer() int {
	a := 100
	defer func() {
		fmt.Println("defer:", a)
		a += 100
	}()
	return a
}
