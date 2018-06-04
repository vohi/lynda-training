package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}

	// custom error object
	myError := errors.New("My error string")
	fmt.Println(myError)

	// attendance := map[string]bool{
	// 	"Ann": true,
	// 	"Mike": true}
}
