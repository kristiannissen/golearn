package main

import "fmt"

type animal struct {
	sound string
	class string
	food  string
}

func (a *animal) eats(s string) {
	a.food = s
}

func main() {
	//
	a := animal{sound: "Puurrr", class: "Mammals"}
	//
	fmt.Println("A", a)
	//
	a.eats("Mice and shit")
	fmt.Println("Eats", a.food)
}
