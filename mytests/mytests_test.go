package mytests

import (
	"testing"
)

func TestMakeSound(t *testing.T) {
	// Simpel test
	got := MakeSound()
	//
	if got != "Hello Kitty" {
		t.Errorf("MakeSound() == ´%s´; want What the fuck!", got)
	}

	// Test table eksempel
	table := []struct {
		input, want string
	}{
		{"Farts", "Only on Mondays"},
		{"Banana", "Fuck NO!"},
		{"Pineapple", "Never on Pizza"},
	}

	for _, v := range table {
		got := Eat(v.input)

		if got != v.want {
			t.Errorf("Eat() got %s, want %s, input %s", got, v.want, v.input)
		}
	}
}
