package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	a := readInput("Acceleration a: ")
	v0 := readInput("Initial Velocity v0: ")
	s0 := readInput("Initial Displacement s0: ")
	fn := GenDisplaceFn(a, v0, s0)
	t := readInput("Time: ")

	fmt.Println("covered distance: ", fn(t))
}

func readInput(msg string) float64 {
	fmt.Print(msg)
	var input string
	fmt.Scanln(&input)

	// ensure right input format
	// https://stackoverflow.com/questions/45686163/how-to-write-isnumeric-function-in-golang/45686455
	if isNumeric(input) == true {
		result, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error: %s", err)
		}
		return result

	} else {
		return readInput(msg)
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}
