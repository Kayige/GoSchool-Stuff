package main

import "fmt"

type newFunction func(int, int) int

var a, b = 1, 2

func main() {
	funcSet := []newFunction{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a - b },
		func(a, b int) int { return a * b },
		func(a, b int) int { return a / b },
	}

	choice := 0
	for {
		fmt.Println("Calculator - 1. Add, 2. Subtract, 3. Multiply, 4. Divide")
		fmt.Println("Enter your choice: ")
		fmt.Scanln(&choice)
		if choice == 1 {
			var fn = funcSet[0]
			fmt.Println("a + b =", fn(a, b))
		} else if choice == 2 {
			var fn = funcSet[1]
			fmt.Println("a - b = ", fn(a, b))
		} else if choice == 3 {
			var fn = funcSet[2]
			fmt.Println("a * b = ", fn(a, b))
		} else if choice == 4 {
			var fn = funcSet[3]
			fmt.Println("a / b = ", fn(a, b))
		} else {
			fmt.Println("Unknown Choice. Exiting")
			break
		}
	}
}
