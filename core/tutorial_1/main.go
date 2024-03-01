package main

import "fmt"

func main() {
	var numerator float32 = 3.14159
	var denominator float32 = 6.9872
	fmt.Printf("numerator: %v, denominator: %v\n", numerator, denominator)
	result, remainder := returnValue(numerator, denominator)
	fmt.Printf("Result: %v \nRemainder: %v \n", result, remainder)
	returnPointer(&numerator, &denominator)
	res := denominator / numerator
	rem := denominator - numerator
	fmt.Printf("point-denominator: %v \n point-numerator: %v \n", denominator, numerator)
	fmt.Printf("RES: %v \n REM: %v \n", res, rem)
}

func returnValue(numerator float32, denominator float32) (float32, float32) {
	remainder := denominator / numerator
	result := denominator - numerator
	return remainder, result
}

func returnPointer(denominator *float32, numerator *float32) {
	*denominator = 8.99
	*numerator = 2.99
}
