package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)
	m["a"] = 0

	v, ok := m["a"]
	if ok {
		fmt.Println(v)
	}

	m["abc"] = 123

	fmt.Println(len(m))

	for k, v := range m {
		fmt.Println(k, v)
	}

	m1 := make(map[string]int)

	go func() {
		for {
			m1["a"] = 1
			time.Sleep(time.Microsecond)
		}
	}()
	go func() {
		for {
			_ = m1["b"]
			time.Sleep(time.Microsecond)
		}
	}()

	select {}

}
