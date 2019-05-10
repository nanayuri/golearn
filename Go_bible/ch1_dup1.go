package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(strings.NewReader("abc\ndsdf\newe\nabc"))
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
