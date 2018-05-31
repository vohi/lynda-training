package main

import "fmt"

func main() {
	doSomething()
	sum := addValues(23, 54)
	fmt.Println("Sum:", sum)
}

func doSomething() {
	fmt.Println("Doing something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}
