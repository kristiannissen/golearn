package mytests

import (
	"strings"
)

func MakeSound() string {
	return "What the fuck!"
	// Output: What the fuck!
}

func Eat(f string) string {
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
