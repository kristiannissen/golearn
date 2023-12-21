package main

import (
	"fmt"
	"strings"
)

var (
	s, b, a string
	f       bool
)

func main() {
	s = "Hello Kitty"
	sc := strings.Clone(s)

	fmt.Println(s == sc)
	// Output: true
	// Fordi s == sc

	// lexicographically sammenligning
	fmt.Println(strings.Compare(s, sc))
	// Output: 0
	// Fordi s == sc

	sc += ", doush!"
	fmt.Println(strings.Compare(s, sc))
	// Output: -1
	// fordi a < b, +1 hvis omvendt

	fmt.Println(strings.Contains(sc, s))
	// Output: true
	// Kan vi finde s i sc?

	fmt.Println(strings.ContainsAny(sc, s))
	// Output: true
	// Fordi Hello eller Kitty findes i sc

	fmt.Println(strings.Count(sc, s))
	// Output: 1
	// Fordi s findes 1 gang i sc

	b, a, f = strings.Cut(sc, "doush")
	fmt.Println(b, a, f)
	// Output: Hello Kitty,  ! true
	// Fordi Hello Kitty, er [b]efore "doush"
	// ! er [a]fter "doush"
	// true fordi vi [f]inder doush i sc

	a, f = strings.CutPrefix(sc, "Hello")
	fmt.Println(a, f)
	// Output:  Kitty, doush! true
	// Fordi Hello_ fjernes
	// true fordi vi finder Hello i sc

	b, f = strings.CutSuffix(sc, "!")
	fmt.Println(b, f)
	// Output: Hello Kitty, doush true
	// Fordi "!" fjernes
	// true fordi ! findes i sc

	fmt.Println(strings.EqualFold(strings.ToUpper(s), strings.ToLower(s)))
	// Output: true
	// Er ens ved simpel UTF-8 case folding

	sp := strings.Split(sc, " ")
	fmt.Println(sp[0], sp[len(sp)-1])
	// Output: Hello doush!
	// Fordi strengen splittes og første og sidste ord skrives
}
