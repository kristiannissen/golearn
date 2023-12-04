package main

import (
	"errors"
	"fmt"
)

// Takes a string argument and returns a string
func say(name string) string {
	return "Hello " + name
}

func sayNoMore(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name is mandatory")
	}
	return "Hello " + name, nil
}

func sayMore(word ...string) {
	s := []string{}
	for _, w := range word {
		s = append(s, w)
	}
	fmt.Println(s)
}

func funk() func(string) string {
	s := "Hello "
	return func(x string) string {
		s = s + x
		return s
	}
}

func main() {
	// Prints Hello Kitty
	fmt.Println(say("Kitty"))
	//
	if _, err := sayNoMore(""); err != nil {
		fmt.Println("Error handling")
	}
	//
	sayMore("Hello", "pretty", "Kitty")
	//
	fmt.Println(funk())
}
