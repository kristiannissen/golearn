package main

import "testing"

func BenchmarkReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadFile()
	}
}
