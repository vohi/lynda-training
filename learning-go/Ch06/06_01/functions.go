package main

import "fmt"

func main() {
	doSomething()
	sum := addValues(23, 54)
	fmt.Println("Sum:", sum)
	sum = addAllValues(12, 54, 79)
	fmt.Println("New Sum:", sum)
}

// lowercase initial character => functions private to the package
func doSomething() {
	fmt.Println("Doing something")
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}
