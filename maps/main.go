// Package maps explores the
// map type
package main

import "fmt"

func main() {
	// Create a new map
	m := map[string]string{} // make(map[string]string)
	// Will print map[]
	fmt.Println(m)
	// Add some stuff to the map
	m["hello"] = "Hello"
	m["pretty"] = "pretty"
	m["kitty"] = "Kitty"
	// Will print map[hello:Hello kitty:Kitty pretty:pretty]
	fmt.Println(m)
	// Len gives us the length of the map
	// Will print Len is 3
	fmt.Println("Len is", len(m))
	// Remove single key
	delete(m, "pretty")
	// Will print map[hello:Hello kitty:Kitty]
	fmt.Println(m)
	// Print key and value paris
	for k, v := range m {
		/**
		 * hello Hello
		 * kitty Kitty
		 */
		fmt.Println(k, v)
	}
	// Error handling
	if _, b := m["pretty"]; b == false {
		// B is false since the key was deleted
		fmt.Println("key does not exist")
	}
	// Clear m
	clear(m)
	// Will print m[]
	fmt.Println(m)
}
