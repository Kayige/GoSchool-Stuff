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

func main() {
	var ProductList []Product
	// struct litereal declaration
	item1 := Product{
		productName:     "Bread",
		productPrice:    1.25,
		productCategory: "Food",
	}
	fmt.Println(item1)

	// Create struct and append it to the slice.
	pl := newProduct("Fish", "Food", 20.0)
	ProductList = append(ProductList, pl)

	// Create 2nd struct and append it to the slice.
	pl = newProduct("Water", "Drink", 0.5)
	ProductList = append(ProductList, pl)

	// Create 3rd struct and append it to the slice.
	pl = newProduct("Detergent", "Household", 10.0)
	ProductList = append(ProductList, pl)

	pl = newProduct("Soap", "Bath", 0.5)
	ProductList = append(ProductList, pl)

	// Show Entire List
	fmt.Println(ProductList)

	// Print First Item on List
	fmt.Println(ProductList[0])

	// Print Second Item on List
	fmt.Println(ProductList[1])

	// Loop over all indexes in the slice.
	// ... Print all struct data.
	for i := range ProductList {
		ProductList := ProductList[i]
		fmt.Printf("\nName: %v \nPrice: $ %v \nCategory: %v\n===========\n", ProductList.productName, ProductList.productPrice, ProductList.productCategory)
	}
}

// function to create item
func newProduct(name string, category string, price float32) Product {
	return Product{
		productName:     name,
		productCategory: category,
		productPrice:    price,
	}
}
