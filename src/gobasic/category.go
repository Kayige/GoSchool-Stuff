package main

import "fmt"

func category() {
	category := []string{"Household", "Food", "Drinks"}

	for i, category := range category {
		fmt.Println(i, category)
	}
}
