package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var bacon string
	bacon = "Bacon ipsum dolor amet chicken bresaola alcatra, biltong t-bone andouille ham hock kevin."
	scanner := bufio.NewScanner(strings.NewReader(bacon))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// Output: bacon, kunne også være os.Stdin

	scanner = bufio.NewScanner(strings.NewReader(bacon))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// Output: bacon
	// ipsum
	// dolor
	// ...

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello,\n")
	fmt.Fprint(w, "Kitty?\n")
	fmt.Println(w.Buffered())
	w.Flush()
	// Output: 14
	// Hello,
	// Kitty?
}
