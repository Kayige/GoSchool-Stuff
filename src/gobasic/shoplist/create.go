package main

import "fmt"

// Variables used for add() & create () set to global
var name, category string
var unitCost float64
var quantity int

func modify() {
	fmt.Println("Which item would you wish to modify?")
	fmt.Scanln(&name)
	fmt.Println("Enter new name. Enter for no change.")
	fmt.Scanln(&name)
	fmt.Println("Enter new Category. Enter for no change.")
	fmt.Scanln(&category)
	fmt.Println("Enter new Quantity. Enter for no change.")
	fmt.Scanln(&unitCost)
	fmt.Println("Enter new Unit Cost. Enter for no change.")
	fmt.Scanln(&quantity)
}

func delete() {
	fmt.Println("Which item do you wish to delete?")
	fmt.Scanln(&name)
	fmt.Println("Item ", &name, "is deleted")
}
