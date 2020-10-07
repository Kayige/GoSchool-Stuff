package main

import (
	"fmt"
	"math"
)

// 1. Write a program that passes in a radius value into a function calCircleArea(), that checks if the radius is more than 0. If it is more than 0, area of the circle, formed from the radius, is calculated. But if radius is less than o, function will return an error.

// 2. Main program to print the area of the circle formed, or error message if no circle is formed.

func calCircleArea(val float64) (float64, error) {
	// check if rad is more than 0
	if val != 0 {
		return math.Pi * val * val, nil
	}
	return 0, fmt.Errorf("radius cannot be 0")
}
func main() {
	radius := 20.0
	area, err := calCircleArea(radius)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Area of Circle:", area)
	}
}
