package main

import (
	"fmt"
	"strconv"
)

func main() {
	var weight, height string
	fmt.Println("Input Weight: ")
	fmt.Scanln(&weight)

	fmt.Println("Input Height: ")
	fmt.Scanln(&height)

	heightValue, _ := strconv.ParseFloat(height, 64)
	weightValue, _ := strconv.ParseFloat(weight, 64)

	bmi := weightValue / (heightValue * heightValue)

	switch {
	case bmi < 18.5:
		fmt.Printf("You are underweight your BMI is: %.2f", bmi)

	case bmi < 24.9:
		fmt.Printf("You are at a Healthy Weight your BMI is: %.2f", bmi)

	case bmi < 29.9:
		fmt.Printf("You are Overweight with BMI: %.2f", bmi)

	case bmi < 34.9:
		fmt.Printf("You are Obese! BMI is: %.2f", bmi)

	case bmi < 39.9:
		fmt.Printf("You are severly obese! BMI: %.2f", bmi)

	case bmi >= 40:
		fmt.Printf("You are Morbidly Obese! BMI: %.2f", bmi)

	default:
		fmt.Println("Invalid Input")

	}
}
