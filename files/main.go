package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var p string
var n int

func main() {
	msg := "Hello Pretty Kitty!"
	// Working Directory
	wd, _ := os.Getwd()
	// File path
	p = filepath.Join(wd, "la-la.txt")
	// Convert string to []byte
	b := []byte(msg)
	// Skriv beskeden i filen og opret filen
	if err := os.WriteFile(p, b, 0666); err != nil {
		// Siger sig selv...
		panic("Fuck!")
	}
	// Opret filen og skriv så til den
	f, err := os.Create(filepath.Join(wd, "lu-lu.txt"))
	if err != nil {
		// PANIC!
		panic("Shit!")
	}
	// Defer lukning af filen
	defer f.Close()
	// Skriv noget i filen
	n, err = f.Write(b)
	fmt.Println("Number of bytes", n)
	// Skriv en streng til filen
	n, err = f.WriteString("Hello Kitty")
	fmt.Println("Number of bytes", n)
	// Sync skrivningen
	f.Sync()
}
