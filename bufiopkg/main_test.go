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

func BenchMarkReadUntil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadUntil(os.Stdout, FILE, byte('\n'))
	}
}

func TestReadUntil(t *testing.T) {
	var o bytes.Buffer

	ReadUntil(&o, FILE, byte('\n'))

	if len(o.String()) == 0 {
		t.Errorf("got len %d", len(o.String()))
	}
}

func TestReadFile(t *testing.T) {

	var o bytes.Buffer

	ReadFile(&o, FILE)

	if len(o.String()) == 0 {
		t.Errorf("got %s", o.String())
	}
}
