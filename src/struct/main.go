// the idea is to store an array of structs
// first create a struct type
// create an array which will store all the struct variables
// print out a list from the array
// from the array print out via index the struct variables.

package main

import "fmt"

type Product struct {
	productName     string
	productPrice    float32
	productCategory string
}

type ProductList []*Product

func main() {

	// struct litereal declaration
	item1 := Product{
		productName:     "Bread",
		productPrice:    1.25,
		productCategory: "Food",
	}

	fmt.Println(item1)

	// Create empty slice of struct pointers.
	ProductList := make([]*Product, 0)

	// Create struct and append it to the slice.
	pl := new(Product)
	pl.productName = "Fish"
	pl.productPrice = 20.0
	pl.productCategory = "Food"

	ProductList = append(ProductList, pl)

	// Create 2nd struct and append it to the slice.
	pl = new(Product)
	pl.productName = "Water"
	pl.productPrice = 0.5
	pl.productCategory = "Drink"

	ProductList = append(ProductList, pl)

	// Create 3rd struct and append it to the slice.
	pl = new(Product)
	pl.productName = "Detergent"
	pl.productPrice = 10.0
	pl.productCategory = "Household"

	ProductList = append(ProductList, pl)

	newProduct("Water", "Drink", 0.50)
	newProduct("Soap", "Bath", 0.5)

	// Loop over all indexes in the slice.
	// ... Print all struct data.
	for i := range ProductList {
		ProductList := ProductList[i]
		fmt.Printf("\nName: %v \nPrice: $ %v \nCategory: %v\n===========\n", ProductList.productName, ProductList.productPrice, ProductList.productCategory)
	}
}

// function to create struct
func (p *Product) newProduct(name string, category string, price float32) {
	pl := new(Product)
	pl.productName = name
	pl.productPrice = price
	pl.productCategory = category
	ProductList = append(ProductList, pl)
}
