package main

import (
	"fmt"
	"os"
)

// Kan ikke ændres
const (
	HOME string = "home"
	USER string = "user"
)

// Kan ændres
var (
	path string
	cmd  string
)

// Init kan bruges til at definere variablerne
func init() {
	fmt.Println("I come first", path, cmd)
	// Output: I come first

	// Definer variabler
	path = os.Getenv("path")
	cmd = "run"
}

func main() {
	// Brug variabler
	fmt.Println("I am second", HOME, USER, path, cmd)
	// Output: I am second home user  run
	// path er ""
}
