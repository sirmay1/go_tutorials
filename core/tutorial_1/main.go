package main
import (
	"fmt"
	"errors"
)

func main() {
	var numerator int = 63
	var denominator int = 0
	var result, remainder, err = intMath(numerator, denominator)
	fmt.Printf("The final result is %v with the remainder of %v", result, remainder)
}
func intMath(numerator int, denominator int) (int, int, error) {
	var err errors
	if denominator == 0 {
		err = errors.New("Sorry! Cannot divide by zero! Please try again sucker!") {
			return 0, 0, err
	}
	var result, remainder int = numerator * denominator
	return result, remainder, err
}
