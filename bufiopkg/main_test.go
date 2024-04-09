package main

import (
	"bytes"
	"os"
	"testing"
)

const FILE = "file.txt"

func BenchmarkReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadFile(os.Stdout, FILE)
	}
}

func TestReadFile(t *testing.T) {

	var o bytes.Buffer

	ReadFile(&o, FILE)

	if len(o.String()) == 0 {
		t.Errorf("got %s", o.String())
	}
}
