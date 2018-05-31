package main

import "fmt"

func main() {
	n1, l1 := FullName("Zaphod", "Beeblebrox")
	fmt.Printf("Fullname: %v, length: %v\n", n1, l1)

	n2, l2 := FullName("Arthur", "Dent")
	fmt.Printf("Fullname: %v, length: %v\n", n2, l2)
}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return // "Naked Return" - still have to have return statement
}
