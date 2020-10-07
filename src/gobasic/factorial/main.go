package main

import (
	"fmt"
	"strconv"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var num string
	fmt.Println("Enter a number: ")
	fmt.Scanln(&num)

	numValue, _ := strconv.ParseInt(num, 10, 0)

	fmt.Println(factorial(int(numValue)))
}
