package main

import "fmt"

func main() {
	var option int
	fmt.Println("Input Temperature Format (1.Kelvin)(2.Celcius) (3.Farenheit):")
	fmt.Scanln(&option)

	var currTemp float32
	fmt.Println("Input Current Temperature: ")
	fmt.Scanln(&currTemp)

	if option == 1 {
		resultFah := currTemp*(5.0/9.0) - 459.67
		resultCelsius := (5.0 / 9.0) * (resultFah - 32)
		fmt.Println("Fahrenheit: ", resultFah, "Celsius: ", resultCelsius)
	} else if option == 2 {
		resultFah := currTemp*(5.0/9.0) + 32
		resultKelvin := (5.0 / 9.0) * (resultFah + 459.65)
		fmt.Println("Fahrenheit: ", resultFah, "Kelvin: ", resultKelvin)
	} else if option == 3 {
		resultKel := (5.0 / 9.0) * (currTemp + 459.67)
		resultCelsius := (5.0 / 9.0) * (currTemp - 32)
		fmt.Println("Kelvin: ", resultKel, "Celsius: ", resultCelsius)
	} else {
		fmt.Println("That is an invalid option")
	}

}
