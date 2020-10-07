package main

import "fmt"

func main() {
	var guess int
	ans := 90
	fmt.Println("Guess the Integer (1-100): ")
	fmt.Scanln(&guess)

	if ans == guess {
		fmt.Println("Number is the Same")
	} else if ans >= guess {
		fmt.Println("TNumber is too Low!")
	} else {
		fmt.Println("Number is too High!")
	}
}
