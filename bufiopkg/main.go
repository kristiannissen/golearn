package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFile() {
	f, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	b := new(strings.Builder)
	if _, err = io.Copy(b, r); err != nil {
		fmt.Println(err)
	}

	fmt.Println(b.String())
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
	ReadFile()

	WriteFile()
}
