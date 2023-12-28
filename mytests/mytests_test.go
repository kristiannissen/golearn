package mytests

import (
	"fmt"
	"os"
	"sort"
	"testing"
)

func setup() {
	// log.Println("Hello from setup()")
}

func teardown() {
	// log.Println("Goodbye from teardown()")
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
	if got != "What the fuck!" {
		t.Errorf("MakeSound() == ´%s´; want What the fuck!", got)
	}
}

func TestFeed(t *testing.T) {

	// Test table eksempel
	table := []struct {
		input, want string
	}{
		{"Farts", "Only on Mondays"},
		{"Banana", "Fuck NO!"},
		{"Pineapple", "Never on Pizza"},
	}

	for _, v := range table {
		got := Feed(v.input)

		if got != v.want {
			t.Errorf("Eat() got %s, want %s, input %s", got, v.want, v.input)
		}
	}
}

func TestFart(t *testing.T) {
	// Alternativ syntax
	if got, want := Fart(), "Fart all day, every day!"; got != want {
		t.Errorf("Fart() got %s, want %s", got, want)
	}
}

// FIXME: Denne fejler!?
func FuzzSortInator(f *testing.F) {
	f.Add([]int{5, 6, 7})

	f.Fuzz(func(t *testing.T, in []int) {
		s := SortInator(in)
		if sort.IntsAreSorted(s) != true {
			t.Fatal("Not sorted", in, s)
		}
	})
}

func BenchmarkFart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fart()
	}
}

func ExampleFart() {
	fmt.Println(Fart())
	// Output: Fart all day, every day
}
