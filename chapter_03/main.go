package main

import "fmt"

// Recursion
func fact(x int) int {
	if x == 1 {
		return x
	} else {
		return x * fact(x-1)
	}
}

func main() {
	result := fact(3)
	fmt.Printf("Factorial of 3 is %d\n", result)
}
