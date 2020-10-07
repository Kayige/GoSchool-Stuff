package main

import "fmt"

const menu = `Shopping List Application
=============================
1. View entire Shopping List
2. Generate List Report
3. Add Items.
4. Modify Items.
5. Delete Items.
Select your Choice:`

func main() {
	shoppingList := make(map[string]ItemInfo)

	for {
		choice := input(menu)

		switch choice {
		case "1":
			viewList(shoppingList)
		case "2":
			report(shoppingList)
		case "3":
			addItem(shoppingList)
		case "4":
			modifyItem(shoppingList)
		case "5":
			deleteItem(shoppingList)
		default:
			fmt.Println("Invalid input")
		}

	}
}
