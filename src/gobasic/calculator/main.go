package main

import "fmt"

func main() {
	var num1, num2, answer float32
	var calc string

	fmt.Println("Input First Number (1-100): ")
	fmt.Scanln(&num1)

	fmt.Println("Input Arithmetic Function (+ - * /): ")
	fmt.Scanln(&calc)

	fmt.Println("Input Second Number (1-100): ")
	fmt.Scanln(&num2)

	if calc == "+" {
		answer = num1 + num2
	}
	if calc == "-" {
		answer = num1 - num2
	}
	if calc == "*" {
		answer = num1 * num2
	}
	if calc == "/" {
		answer = num1 / num2
	}

	fmt.Println(answer)
}
