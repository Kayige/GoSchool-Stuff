package main

import (
	"fmt"
)

type currency struct {
	currencyName string
	rate         float64
}

var m map[string]currency

func main() {

	m = make(map[string]currency)

	m["USD"] = currency{"US Dollar", 1.1318}
	m["JPY"] = currency{"Japanese YEN", 121.05}
	m["GBP"] = currency{"Pounds Sterling", 0.90630}
	m["CNY"] = currency{"Chinese Yuan Renminbi", 7.9944}
	m["SGD"] = currency{"Singapore Dollar", 1.5743}
	m["MYR"] = currency{"Malaysian Ringgit", 4.8390}

	fmt.Println(m["USD"].currencyName, "-", m["USD"].rate)
	fmt.Println(m)

	var currencyFrom string
	fmt.Println("Enter your currency:")
	fmt.Scanln(&currencyFrom)

	var currencyAmt float64
	fmt.Println("Enter Amount: ")
	fmt.Scanln(&currencyAmt)

	var currencyTo string
	fmt.Println("Enter currency convert to: ")
	fmt.Scanln(&currencyTo)

	result := currencyAmt / m[currencyFrom].rate * m[currencyTo].rate
	fmt.Println("Currency", currencyFrom, "convert to", currencyTo, result)
}
