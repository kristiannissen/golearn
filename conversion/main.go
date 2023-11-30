// Conversion explores types and conversion
// you can find more at https://go.dev/ref/spec#Conversions
package main

import "fmt"

func main() {
	f := 42.0
	// Prints F is of type float64
	fmt.Printf("F is of type %T \n", f)
	/*
	 * Now let's convert it to a float32
	 */
	// Prints F is of type float32
	fmt.Printf("F is of type %T\n", float32(f))
	/*
	 * Play with byte and string
	 */
	b := []byte{'a', 'b'}
	// Will print [97 98]
	fmt.Println(b)
	// Convert to string
	// Will print ab
	fmt.Println(string(b))
	// More byte
	c := []byte("")
	// Will print []
	fmt.Println(c)
	// Will print [104 101 108 108 111 32 107 105 116 116 121]
	d := []byte("hello kitty")
	fmt.Println(string(d))
	// Will print ""
	fmt.Print(d)
}
