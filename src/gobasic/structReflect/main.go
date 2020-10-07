package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	fName        string
	lName        string
	uID          int
	InvoiceTotal float64
}

func inspect(n interface{}) {
	refType := reflect.TypeOf(n)
	refValue := reflect.ValueOf(n)

	fmt.Println("Num of fields: ", refType.NumField())
	for i := 0; i < refType.NumField(); i++ {
		fmt.Println(refType.Field(i).Name, "value: ", refValue.Field(i), "Type: ", refType.Field(i).Type)
	}
}

func main() {

	customer1 := customer{
		fName:        "John",
		lName:        "Wick",
		uID:          123509612,
		InvoiceTotal: 5.3462123,
	}
	inspect(customer1)
}
