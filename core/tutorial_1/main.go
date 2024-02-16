package main

import (
	"fmt"
)

func main() {
	var numeric int = 999
	var increment int = 99
	var remainder, results = intMath(numeric, increment)
	fmt.Printf("This is your remainder: %v and here is your result: %v", remainder, results)
}
func intMath(numeric int, increment int) (int, int) {
	var remainder int = numeric - increment
	var results int = numeric + increment
	return remainder, results
}
