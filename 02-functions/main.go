package main

import (
	"errors"
	"fmt"
)

func main() {
	// ******** Functions ***********

	// Simple function
	someVal := "Hello, World!"
	printMe(someVal) // calling the function

	// Functions with parameters and return values
	numerator, denominator := 20, 0
	var division, remainder, err = intDivision(numerator, denominator) // calling the function

	// handling errors, it's general design pattern in Go to return an error as the last return value of a function, and the caller is responsible for checking if the error is nil before using the other return values.
	if err != nil {
		fmt.Println("Error:", err.Error()) // Error: Cannot divide by zero
		return
	}

	fmt.Printf("Result is %d and remainder is %d\n", division, remainder) // 6 and 2
}

func printMe(value string) {
	fmt.Println(value)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error // default is nil
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err // return default values in case of error
	}

	res := numerator / denominator
	remainder := numerator % denominator
	return res, remainder, err
}
