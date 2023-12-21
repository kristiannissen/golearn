package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("vim-go", n.Local())
	// Output: 2023-12-21 15:07:45.636693015 +0100 CET

	f := n.Add(time.Hour * 24)
	fmt.Println(f)
	// Output: 2023-12-22 15:09:11.400342555 +0100 CET m=+86400.000036361

	f = n.AddDate(0, 0, 365)
	fmt.Println(f)
	// Output: 2024-12-20 15:11:13.557246166 +0100 CET

	fmt.Println(f.After(n))
	// Output: true

	fmt.Println(n.Before(f))
	// Output: true

	fmt.Println(n.Clock())
	// Output: 15 15 26

	fmt.Println(n.Weekday())
	// Output: Thursday

	fmt.Println(n.Format("Mon 02 2006 -07:00"))
	// Output: Thu 21 2023 +01:00

	d := "22. Dec 2023 20:00"
	fmt.Println(time.Parse("02. Jan 2006 15:04", d))
	// Output: 2023-12-22 20:00:00 +0000 UTC <nil>
}
