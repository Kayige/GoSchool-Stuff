package main

import (
	"fmt"
	"strconv"
)

func deleteItem(shoppingList ShoppingList) {
	fmt.Println("\nDelete Item")
	name := input("Enter item name to delete:")
	if flag := shoppingList.deleteItem(name); flag {
		fmt.Printf("[Deleted %s]\n", name)
	} else {
		fmt.Println("Item not found. Nothing to delete!")
	}
}

func modifyItem(shoppingList ShoppingList) {
	fmt.Println("\nModify Item")
	name := input("Which item would you wish to modify?")
	item, flag := shoppingList.getItem(name)
	if !flag {
		fmt.Println("[Item not in list]")
		return
	}

	fmt.Printf("Current item name is %s - Category is %s - Quantity is %d - Unit Cost %.2f\n", name, item.Category, item.Quantity, item.UnitCost)

	n := input("Enter new name. Enter for no change")
	if n == name {
		n = ""
	}
	if _, flag := shoppingList.getItem(n); flag {
		fmt.Println("[this name is already taken by another item]")
		return
	}
	if n == "" {
		n = name
		defer fmt.Println("No changes to item made.")
	}

	c := input("Enter new Category. Enter for no change")
	if c != "" {
		item.Category = c
	} else {
		defer fmt.Println("No changes to category made.")
	}

	q := input("Enter new Quantity. Enter for no change")
	num, err := strconv.ParseInt(q, 10, 0)
	if q != "" {
		if err != nil {
			fmt.Println("invalid input")
			return
		}
		item.Quantity = int(num)
	} else {
		defer fmt.Println("No changes to quantity made.")
	}

	nC := input("Enter new Unit cost. Enter for no change")
	cost, err := strconv.ParseFloat(nC, 64)
	if nC != "" {
		if err != nil {
			fmt.Println("invalid input")
			return
		}
		item.UnitCost = cost
	} else {
		defer fmt.Println("No changes to unit cost made.")
	}

	shoppingList.modifyItem(name, n, item)
}

func addItem(shoppingList ShoppingList) {
	fmt.Println("\nAdd Item")
	name := input("What is the name of your item?")
	if _, flag := shoppingList.getItem(name); flag {
		fmt.Println("Item already in the list.")
		return
	}
	category := input("What category does it belong to?")
	quantity := input("How many units are there?")
	unitCost := input("How much does it cost per unit?")

	if checkEmpty(name, category, quantity, unitCost) {
		fmt.Println("invalid input")
	} else if quantity, err := strconv.ParseInt(quantity, 10, 0); err != nil {
		fmt.Println("invalid input")
	} else if unitCost, err := strconv.ParseFloat(unitCost, 64); err != nil {
		fmt.Println("invalid input")
	} else {
		shoppingList.addItem(name, ItemInfo{Category: category, Quantity: int(quantity), UnitCost: unitCost})
		fmt.Printf("[Added %s]\n", name)
	}
}

func viewList(shoppingList ShoppingList) {
	fmt.Println("\nShopping List Contents:")
	if len(shoppingList) == 0 {
		return
	}
	for name, item := range shoppingList {
		fmt.Println(itemString(name, item))
	}
}
