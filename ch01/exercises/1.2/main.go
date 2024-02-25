package main

import (
	"fmt"
	"os"
)

// Modify the echo program to print the index and value of each of its arguments, one per line.

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
