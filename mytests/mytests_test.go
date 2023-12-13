package mytests

import (
	"log"
	"os"
	"testing"
)

func setup() {
	log.Println("Hello from setup()")
}

func teardown() {
	log.Println("Goodbye from teardown()")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
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
