package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fileLines := make(map[string]map[string]bool)

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileLines, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileLines, arg)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for file := range fileLines[line] {
				fmt.Printf("\t%s\n", file)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileLines map[string]map[string]bool, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if counts[line] > 1 {
			if fileLines[line] == nil {
				fileLines[line] = make(map[string]bool)
			}
			fileLines[line][fileName] = true
		}
	}
}
