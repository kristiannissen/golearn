package main

import (
	"fmt"
	"sort"
)

func main() {
	fruits := []struct {
		Name string
		qty  int
	}{
		{"Apple", 4},
		{"Banana", 6},
		{"Shit", 8},
		{"Stuff", 8},
	}

	sort.Slice(fruits, func(i, j int) bool {
		return fruits[i].qty < fruits[j].qty
	})
	fmt.Println(fruits)
	// Output: [{Apple 4} {Banana 6} {Shit 8} {Stuff 8}]

	sort.SliceStable(fruits, func(i, j int) bool {
		return fruits[i].qty < fruits[j].qty
	})
	fmt.Println(fruits)
	// Output: [{Apple 4} {Banana 6} {Shit 8} {Stuff 8}]

	numbers := []int{5, 2, 4, 7, 8, 9, 1}
	fmt.Println(numbers)
	// Output: [5 2 4 7 8 9 1]

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println(numbers)
	// Output: [1 2 4 5 7 8 9]
}
