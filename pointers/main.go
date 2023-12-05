package main

import "fmt"

type Animal struct {
	kind   string
	number int
}

func (a Animal) say() string {
	return "The animal says what?!"
}

func (a *Animal) howMany() int {
	a.number++
	return a.number
}

var list []string

func listLen(l []string) int {
	return len(l)
}

func addItem(s string, l *[]string) {
	*l = append(*l, s)
}

func main() {
	// Printer 0
	fmt.Println(listLen(list))
	//
	addItem("Hello", &list)
	// Printer 1
	fmt.Println(listLen(list))
	//
	animal := Animal{"Cat", 0}
	// Printer Animal main.Animal
	fmt.Printf("Animal %T\n", animal)
	// Printer The animal says what?!
	fmt.Println(animal.say())
	// Printer 1
	fmt.Println(animal.howMany())
	// Printer 2
	fmt.Println(animal.howMany())
	// Printer 3
	fmt.Println(animal.howMany())

}
