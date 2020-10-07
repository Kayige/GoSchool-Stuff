package main

import "fmt"

func main() {
	num1 := 17
	num2 := 24
	if num1 == num2 {
		fmt.Println("Both are equal")
		fmt.Println("The difference is 0")
	} else if num1 > num2 {
		diff := num1 - num2
		fmt.Println("Num1 is Bigger than Num2")
		fmt.Println("The difference is ", diff)
	} else if num2 > num1 {
		diff := num2 - num1
		fmt.Println("Num2 is Bigger than Num1")
		fmt.Println("The difference is ", diff)
	} else {
		fmt.Println("Invalid Numbers")
	}
}
