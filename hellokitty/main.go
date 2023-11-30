// Package HelloKitty explores the fmt package from the std
// look at https://pkg.go.dev/fmt@go1.21.4
package main

import "fmt"

func main() {
	// Print a line of text
	// Prints Hello Kitty
	fmt.Println("Hello Kitty!")
	// Print type using %T
	var name string
	name = "Kitty"
	// Prints Name is string
	fmt.Printf("Name is %T\n", name)
}
