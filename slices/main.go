// The slices packages explores
// different ways to use slices
package main

import "fmt"

func main() {
	/*
	 * Create a new slices of type string
	 */
	var s []string
	// Will print S lenght is 0
	fmt.Printf("S lenght is %d \n", len(s))
	// Let us add something to s
	s = append(s, "Hello")
	s = append(s, "pretty")
	s = append(s, "Kitty")
	// Will print S length is 2
	fmt.Printf("S length is %d \n", len(s))
	// Let's print part of s
	// Will print [pretty Kitty]
	fmt.Println("", s[1:])
	// Will print [Hello]
	fmt.Println("", s[0:1])
	// Will print pretty
	fmt.Println("", s[1])
	// Loop over slice
	for i := 0; i < len(s); i++ {
		/*
		 * Hello
		 * pretty
		 * Kitty
		 */
		fmt.Println(s[i])
	}
	for i, v := range s {
		/*
		 * 0 Hello
		 * 1 pretty
		 * 2 Kitty
		 */
		fmt.Println(i, v)
	}
}
