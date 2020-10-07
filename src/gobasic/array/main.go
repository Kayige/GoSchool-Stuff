package main

import "fmt"

func main() {
	var a = []string{"Week 1 - $9.50", "Week 2 - $8.00", "Week 3 - $10.20", "Week 4 - $7.40"}
	fmt.Println("Operating Systems List.")
	fmt.Println(len(a))
	fmt.Println(cap(a))

	fmt.Println(a[2])
	a[2] = "Week 3 - $9.80"
	fmt.Println(a[2])

	a = append(a, "Week 5 - $8.40", "Week 6 - $9.40", "Week 7 - $7.20")
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	b := a[2:8]
	fmt.Println(b)
}
