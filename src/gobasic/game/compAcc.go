package main

import "fmt"

type compAcc struct {
	title string
	price float32
}

func (c compAcc) print() {
	fmt.Println(c.title, c.price)
}
