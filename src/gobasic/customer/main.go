package main

import "fmt"

func main() {
	var customer1 = customer{
		fName:   "Michael",
		lName:   "Jordan",
		uName:   "M32020",
		pw:      "1234567",
		email:   "mj@email.com",
		phone:   7654321,
		address: "18227 Captain Greens Goad Cornelius, NC 20031",
	}

	customer1.printAllInfo()
	fmt.Println(customer1.printUserAddress())
	fmt.Println(customer1.printUserCredentials())

}
