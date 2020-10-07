package main

import (
	"fmt"
	"strconv"
)

func report() {
	for {
		var choice string
		fmt.Println(`
Generate Report
1. Total Cost of each category.
2. List of item by category.
3. Main Menu.
Choose your report: 
					`)
		fmt.Scanln(&choice)

		choiceValue, _ := strconv.Atoi(choice)

		if choiceValue == 1 {
			fmt.Println("Total cost by Category.")
		} else if choiceValue == 2 {
			fmt.Println("List by Category.")
		} else if choiceValue == 3 {
			break
		} else {
			fmt.Println("Invalid Input")
			continue
		}
	}
}
