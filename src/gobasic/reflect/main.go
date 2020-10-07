package main

import (
	"fmt"
	"reflect"
)

func inspect(n interface{}) {
	refType := reflect.TypeOf(n)
	fmt.Println("Content ", n, "Name: ", refType.Name(), "Kind: ", refType.Kind())
}
func main() {
	str := "This is a string"
	inspect(str)
	num := 12345
	inspect(num)
	flt := 1.2345
	inspect(flt)
	bool := true
	inspect(bool)
}
