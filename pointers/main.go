package main

import "fmt"

var list []string

func listLen(l []string) int {
	return len(l)
}

func addItem(s string, l *[]string) {
	*l = append(*l, s)
}

func main() {
	fmt.Println(listLen(list))
	//
	addItem("Hello", &list)
	//
	fmt.Println(listLen(list))
}
