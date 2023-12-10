package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 0 {
		// Prints /tmp/go-build3684091303/b001/exe/main, Hello, Kitty
		// os.Args[0] er program navnet
		fmt.Println(strings.Join(os.Args, ", "))
	}

	var name string
	flag.StringVar(&name, "name", "Pussy", "The name")

	flag.Parse()

	// go run main.go -name=Kitty Prints: Name Kitty
	// go run main.go Prints: Name "Pussy"
	fmt.Println("Name", name)
	// Vis alle
	fmt.Println(flag.Args())
}
