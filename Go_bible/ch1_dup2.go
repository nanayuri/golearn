package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, strings.Split(line, "|")[1], strings.Split(line, "|")[0])
		}
	}
}

func countLines(f *os.File, counts map[string]int, file_name string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()+"|"+file_name]++
	}
}
