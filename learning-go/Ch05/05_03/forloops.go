package main

import "fmt"

func main() {

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := range colors {
		fmt.Println(i)
	}

	for sum < 1000 { // while loop
		sum += sum
		if sum > 900 {
			goto endofprogram
		}
	}
	fmt.Println(sum)

endofprogram:
	fmt.Println("The end")
}
