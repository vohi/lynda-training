package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	result := ""
	switch dow := rand.Intn(6) + 1; dow {
	case 6:
		result = "Saturday"
	case 7:
		result = "Sunday"
	default:
		result = "Weekday"
	}

	fmt.Println("It's a", result)

	x := -42

	switch {
	case x < 0:
		result = "Less than zero"
		// fallthrough // optional fallthrough avoids need for break
	case x == 0:
		result = "Equal zero"
	default:
		result = "Greater than zero"
	}
	fmt.Println("It's a", result)
}
