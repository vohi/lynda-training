package main

import (
	"fmt"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)]) // len(colors) would be the default
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1]) // 0 is the default
	fmt.Println(colors)

	numbers := make([]int, 5, 5) // type, size, capacity (defaults to size)
	numbers[0] = 1
	numbers[1] = 7
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156
	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println(cap(numbers)) // prints 10 - go's growth strategy for slices
}
