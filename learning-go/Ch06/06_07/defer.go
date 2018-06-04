package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Close the file")
	fmt.Println("Open the file")

	// add function call to execution stack - LIFO
	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")

	myFunc()

	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")
	fmt.Println("Undeferred")

	x := 1000
	defer fmt.Println("Value of x:", x)
	x++
	fmt.Println("Value of x++:", x)
}

func myFunc() {
	defer fmt.Println("Deferred in the function")
	fmt.Println("Not deferred in the function")
}
