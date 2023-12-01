// The statement package explores
// https://go.dev/ref/spec#Statements
package main

import "fmt"

func main() {
	// Simple for loop
	// will print
	// I is  0
	// I is  1
	// I is  2
	for i := 0; i < 3; i++ {
		fmt.Println("I is ", i)
	}
	// Switch
	// will print Oh
	i := 1
	switch i {
	case 0:
		fmt.Println("No no")
	case 1:
		fmt.Println("Oh ")
	default:
		fmt.Println("no!")
	}
	// If statement
	// will print J is less than K
	j := 2
	if k := isFour(); j < k {
		fmt.Println("J is less than K")
	} else {
		fmt.Println("Programmer skrewed up!")
	}
	// Will print Hello 2
	l, err := fmt.Printf("Hello %d \n", j)
	// Often used syntax
	if err != nil {
		fmt.Errorf("Err %s", err)
	}
	// Will print L is 9, the lenght of the bytes
	// written
	fmt.Printf("L is %d \n", l)
}

func isFour() int {
	return 4
}
