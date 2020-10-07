package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var num string
		fmt.Println("Enter a number: ")
		fmt.Scanln(&num)

		numValue, _ := strconv.ParseInt(num, 10, 0)

		if numValue%2 == 0 {
			fmt.Println("The number", numValue, "is Even")
		} else {
			fmt.Println("The number", numValue, "is Odd.")
		}

		if numValue >= 10 {
			fmt.Println("The Number", numValue, "has 2 digits.")
		} else {
			fmt.Println("The Number", numValue, "has 1 digit.")
		}
	}
}
