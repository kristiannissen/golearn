package main

import (
	"fmt"
	"log"
	"os"
)

type animal struct {
	sound string
	class string
	food  string
}

func (a *animal) eats(s string) {
	a.food = s
}

// Embedding
type RingerStringer struct {
	Inp string
	// RingerStringer kan nu bruge RingerStringer.Println
	*log.Logger
}

func main() {
	//
	a := animal{sound: "Puurrr", class: "Mammals"}
	//
	fmt.Println("A", a)
	//
	a.eats("Mice and shit")
	fmt.Println("Eats", a.food)

	// RingerStringer
	r := &RingerStringer{"Hello", log.New(os.Stderr, "Ringer: ", log.Ldate)}
	r.Printf("Hello Ringer Stringer %s", r.Inp)
	// Output: Ringer: 2024/01/02 Hello Ringer Stringer Hello
}
