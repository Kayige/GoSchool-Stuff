package main

import (
	"fmt"
)

var (
	message, firstName, lastName, account, password string
	weight, height, credit                          float32
	age                                             int
	subscribedUser                                  bool
)

func main() {
	message = "This is the account information."
	firstName = "Luke"
	lastName = "Skywalker"
	age = 20
	weight = 72.0
	height = 1.72
	credit = 123.55
	account = "admin"
	password = "password"
	subscribedUser = true
	fmt.Printf("%T %T %T %T %T %T %T %T %T %T\n", message, firstName, lastName, age, weight, height, credit, account, password, subscribedUser)
	fmt.Println(message, firstName, lastName, age, weight, height, credit, account, password, subscribedUser)
}
