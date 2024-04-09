package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadScanner() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println("fmt:", s.Text())
	}
	if err = s.Err(); err != nil {
		log.Fatal(err)
	}
}

func ReadUntil() {
	f, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	d, e := r.ReadString('\n')
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(d)
}

// ReadFile uses bufio to read the content of the file passed as n
// the content is written to w. W needs to implement https://pkg.go.dev/io#Writer
// interface. The file content is written as a string to w
func ReadFile(w io.Writer, n string) {
	f, err := os.Open(n)
	if err != nil {
		// Print error and exit
		log.Fatalf("ReadFile: %s", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	b := new(strings.Builder)
	if _, err = io.Copy(b, r); err != nil {
		// Print error and exit
		log.Fatalf("ReadFile: %s", err)
	}

	fmt.Fprintf(w, b.String())
}

func WriteFile() {
	f, err := os.Create("hello_kitty.txt")
	if err != nil {
		fmt.Println(err)
	}

	w := bufio.NewWriter(f)
	_, err = w.Write([]byte("Hello Kitty"))
	if err != nil {
		fmt.Println("1", err)
	}

	if err = w.Flush(); err != nil {
		fmt.Println("2", err)
	}
}

func main() {
	ReadFile(os.Stdout, "file.txt")

	// WriteFile()

	// ReadUntil()

	// ReadScanner()
}
