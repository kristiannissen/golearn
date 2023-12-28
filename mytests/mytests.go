package mytests

import (
	"sort"
	"strings"
)

func MakeSound() string {
	return "What the fuck!"
}

func Feed(f string) string {
	var resp string

	switch strings.ToLower(f) {
	case "banana":
		resp = "Fuck NO!"
	case "pineapple":
		resp = "Never on Pizza"
	default:
		resp = "Only on Mondays"
	}

	return resp
}

func Fart() string {
	return "Fart all day, every day!"
}

func SortInator(num []int) []int {
	sort.Slice(num, func(i, j int) bool {
		return num[i] < num[j]
	})
	return num
}
