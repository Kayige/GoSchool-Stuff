// the idea is to store an array of structs
// first create a struct type
// create an array which will store all the struct variables
// print out a list from the array
// from the array print out via index the struct variables.

package main

import "fmt"

type product struct {
	productName     string
	productPrice    float32
	productCategory string
}

var productList *product

func main() {

	item1 := Product{
		productName:     "Bread",
		productPrice:    1.25,
		productCategory: "Food",
	}

	fmt.Println(item1)
	fmt.Println(ProductList)
}

func newProduct(name string, category string, price float32) *Product{
	pl := ProductList
	if pl == nil {
		return pl[0] := &Product{
			productName: name,
			productPrice: price,
			productCategory: category,
		}
		
	} else {
		for i=0; i < len(pl); i++ {
			if pl[i] == nil {
				return pl[i] := &Product{
					productName: name,
					productPrice: price,
					productCategory: category,
				}
			} 
		}
	}
	

	}
}