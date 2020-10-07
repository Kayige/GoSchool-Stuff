package main

import (
	"fmt"
)

func input(message string) (answer string) {
	fmt.Println(message)
	fmt.Scanln(&answer)
	return
}

func checkEmpty(check ...string) bool {
	for _, string := range check {
		if string == "" {
			return true
		}
	}
	return false
}

func itemString(name string, item ItemInfo) string {
	return fmt.Sprintf("Category: %s - Item: %s Quantity: %d Unit Cost: %.2f", item.Category, name, item.Quantity, item.UnitCost)
}

const message = `Generate Report
			1. Total Cost of each category.
			2. List of item by category.
			3. Main Menu.
			Choose your report: `

func report(shoppingList ShoppingList) {
	if len(shoppingList) == 0 {
		fmt.Println("[The list is Empty]")
		return
	}

	costByCat := map[string]float64{}
	itemByCat := map[string][]string{}

	for name, item := range shoppingList {
		costByCat[item.Category] = costByCat[item.Category] + (float64(item.Quantity) * item.UnitCost)

		itemByCat[item.Category] = append(itemByCat[item.Category], itemString(name, item))
	}

	choice := input(message)

	switch choice {
	case "1":
		fmt.Println("\nTotal cost by Category")
		for category, total := range costByCat {
			fmt.Printf("%s cost : %.2f\n", category, total)
		}
	case "2":
		fmt.Println("\nList by Category")
		for _, items := range itemByCat {
			for _, item := range items {
				fmt.Println(item)
			}
		}
	case "3":
		return
	default:
		fmt.Println("Invalid input.")
	}
}
