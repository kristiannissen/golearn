package main

import "fmt"

// Takes a string argument and returns a string
func say(name string) string {
	s := "Hello " + name
	return s
}

func main() {
	// Prints Hello Kitty
	fmt.Println(say("Kitty"))
}
