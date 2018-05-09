package main

import "fmt"

func main() {

	// var x float64 = 42
	var result string

	if x := 42; x < 0 {
		result = "less than zero"
	} else if x == 0 {
		result = "equal zero"
	} else {
		result = "greater than zero"
	}

	fmt.Println("x is", result)
}
