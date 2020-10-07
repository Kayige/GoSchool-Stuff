package main

import "fmt"

func main() {
	result := sum(10, 5)
	fmt.Println(result)
}

func sum(num1, num2 int) int {
	result := num1 + num2
	return result

}
