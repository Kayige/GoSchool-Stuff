package main

import "fmt"

type customer struct {
	fName   string
	lName   string
	uName   string
	pw      string
	email   string
	phone   int
	address string
}

func (c customer) printUserCredentials() (string, string) {
	return c.uName, c.pw
}

func (c customer) printUserAddress() string {
	return c.address
}

func (c customer) printAllInfo() {
	fmt.Println(c.fName, c.lName, c.uName, c.pw, c.email, c.phone, c.address)
}
