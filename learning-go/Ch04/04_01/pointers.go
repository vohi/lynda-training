package main

import "fmt"

func print(p *int) {
	if p != nil {
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}
}

func main() {
	var p *int
	print(p)

	var v int = 42
	p = &v
	print(p)

	var value1 float64 = 42.13
	pointer1 := &value1

	*pointer1 = *pointer1 / 31
	fmt.Println("Value1:", *pointer1)
	fmt.Println("Value1:", value1)
}
