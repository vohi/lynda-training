package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	fmt.Printf("Map: %v (%v)\n", m, len(m))
	if m != nil {
		m["fourty-two"] = 42
	}

	var newMap = new(map[string]int) // returns pointer
	if newMap != nil {
		// (*newMap)["fourty-two"] = 42 // this panics
		*newMap = make(map[string]int)
		(*newMap)["fourty-two"] = 42
	}
	fmt.Printf("newMap: %v (%v)\n", *newMap, len(*newMap))

	var makeMap = make(map[string]int, 0)
	if makeMap != nil {
		makeMap["fourty-two"] = 42
	}
	fmt.Printf("makeMap: %v (%v)\n", makeMap, len(makeMap))
}
