package main

import (
	"fmt"
	"time"
)

func main() {
	// package level functions ~ static methods
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)

	// type methods to operate on objects
	fmt.Println("The month is", t.Month())
	fmt.Println("The day is", t.Day())
	fmt.Println("The weekday is", t.Weekday())

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	// reference date is 1-2-3-4-5-6-7: January 2nd, 03:45; 2006; GMT-7
	longFormat := "Monday, January 2, 2006"
	shortFormat := "2.1.06" // day-month-year
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))
}
