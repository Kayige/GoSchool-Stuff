package main

import "fmt"

func main() {

	for {
		num := 0
		fmt.Println("Input Number: ")
		fmt.Scanln(&num)
		if num%2 == 0 {
			fmt.Println("Number is Even")
		} else {
			fmt.Println("Number if Odd")
		}
	}
}
