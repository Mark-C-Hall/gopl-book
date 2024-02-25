package main

import "testing"

func BenchmarkInefficient(b *testing.B) {
	words := []string{"one", "two", "three", "four", "five"}
	for i := 0; i < b.N; i++ {
		inefficient(words)
	}
}

func BenchmarkEfficient(b *testing.B) {
	words := []string{"one", "two", "three", "four", "five"}
	for i := 0; i < b.N; i++ {
		efficient(words)
	}
}
