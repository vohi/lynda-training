package main

import (
	"fmt"
	"time"
)

func main() {
	// package level functions ~ static methods
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)
}
