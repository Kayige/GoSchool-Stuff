package main

import "fmt"

// Given the length of all three sides of a triangle, a, b, c, the area of the triangle can be computed by using Heron’s formula:
// area = √(𝑠(𝑠−𝑎)(𝑠−𝑏)(𝑠−𝑐)) where s = (𝑎+𝑏+𝑐)/2
// The values of the three sides can form a triangle only if sum of any two sides is larger than the third side.
// 1. Define an invalidTriangleError error type that consists of the 3 sides (that gives the invalid triangle) and an error message for the error.
// 2. Define an invalidSideError error type that consists of the erroneous side value and error message.
// 3. Write a createTriangle() function that checks if any of the 3 sides is less than 0. If yes, an invalidSideError instance is constructed and returned. It then checks if the 3 sides can form a triangle, in which case if no, an invalidTriangleError would be returned instead. If none of the errors occur, the function will go ahead and calculate the area of the triangle formed and return the calculated value.
// 4. Write a main program that will test the above.
// (Hint can test with values 20, 5, 12 which will not form a triangle, 20, 5, 20 will form a triangle)

// TriangleType is the alias of type int that represents the type of the triangles
type TriangleType int

/*
* None: There is an error such as the float64 overflow or it is not a triangle
* Scalene: No sides of the triangle are equal
* Isosceles: Any two sides of the triangle are equal
* Equilateral: All sides of the triangle are equal
 */
const (
	None = TriangleType(iota)
	Scalene
	Isosceles
	Equilateral
)

// DetermineTriangleType is used to determine the type by the input
// It will first check if it is a triangle
// Then it will check the type.
func DetermineTriangleType(a float64, b float64, c float64) TriangleType {
	if !IsTriangle(a, b, c) {
		return None
	}
	if AllSidesAreEqual(a, b, c) {
		return Equilateral
	} else if TwoSidesAreEqual(a, b, c) {
		return Isosceles
	}
	return Scalene
}

// IsTriangle is used to determine if it is a triangle by the input.
// It will first check if all the sides are positive numbers.
// Then it will check if the sum of any two sides is greater than the third side.
func IsTriangle(a float64, b float64, c float64) bool {
	if AllSidesArePositive(a, b, c) {
		return AnyTwoSidesAreGreaterThanTheThrid(a, b, c)
	}
	return false
}

// AllSidesArePositive is used to check if all the sides are positive numbers
func AllSidesArePositive(a float64, b float64, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	return true
}

// AnyTwoSidesAreGreaterThanTheThrid is used to check if the sum of any two sides is greater than the third side.
func AnyTwoSidesAreGreaterThanTheThrid(a float64, b float64, c float64) bool {
	// We may need to consider about the float64 overflow if a + b > the max value of float64
	if a+b <= c || a+c <= b || b+c <= a {
		return false
	}
	return true
}

// AllSidesAreEqual is used to check if all the sides are equal or not.
func AllSidesAreEqual(a float64, b float64, c float64) bool {
	if a == b && b == c {
		return true
	}
	return false
}

// TwoSidesAreEqual is used to check if any two sides are equal or not.
func TwoSidesAreEqual(a float64, b float64, c float64) bool {
	if a == b || a == c || b == c {
		return true
	}
	return false
}

func main() {
	fmt.Println(DetermineTriangleType(20, 5, 12))
	fmt.Println(DetermineTriangleType(20, 5, 20))
	fmt.Println(DetermineTriangleType(1, 2, 4))
	fmt.Println(DetermineTriangleType(3, 4, 5))
	fmt.Println(DetermineTriangleType(4, 4, 5))
	fmt.Println(DetermineTriangleType(5, 5, 5))
}
