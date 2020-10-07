package main

import (
	"fmt"
	"strconv"
)

type itemData struct {
	itemCategory *[]string
	itemQuantity int
	unitCost     float64
}

type itemMap map[string]*itemData

// func (i iteminfo) print() {
// 	fmt.Println("Category: ", i.category, "Quantity: ", i.quantity, "Unit Cost: ", i.unitcost)
// }

func main() {
	var category itemData
	category.itemCategory = &[]string{"Food", "Drink", "Household"}

	im := make(itemMap)

	for { // loop for menu option.
		var choice string
		fmt.Println(`
Shopping List Application
=============================
1. View entire Shopping List
2. Generate List Report
3. Add Items.
4. Modify Items.
5. Delete Items.
Select your Choice:
				 `)

		fmt.Scanln(&choice)
		choiceValue, _ := strconv.Atoi(choice)

		// var shoplist list
		// shoplist = append(shoplist)

		// testing input
		im["Bread"] = &itemData{itemCategory: &[]string{"Food"}, itemQuantity: 3, unitCost: 5.5}

		// print input
		for key := range im {
			fmt.Println("Category", im[key].itemCategory, "- Item:", key, "Quantity:", im[key].itemQuantity, "Cost:", im[key].unitCost)
		}

		// // Category Slice
		// category := []string{"Household", "Food", "Drinks"}

		// for i, category := range category {
		// 	fmt.Println(i, category)
		// }

		switch choiceValue {
		case 1:
			// shoplist.print()
		case 2:
			report() // runs report options from report.go
		case 3:
			var name, category string
			var unitCost float64
			var quantity int
			fmt.Println("What is the name of your item?")
			fmt.Scanln(&name)
			fmt.Println("What category does it belong to?")
			fmt.Scanln(&category)
			fmt.Println("How much does it cost per unit?")
			fmt.Scanln(&unitCost)
			fmt.Println("How many quantity are there?")
			fmt.Scanln(&quantity)

			im[name] = &itemData{itemCategory: &[]string{category}, itemQuantity: quantity, unitCost: unitCost}
		case 4:
			modify() // runs modify from create.go
		case 5:
			delete() // delete from create.go
		default:
			fmt.Println("Invalid Input")
			continue
		}
	}
}
