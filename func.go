package main

import "fmt"

// Function with parameters and return value
func add(x, y int) int {
	return x + y
}

func main() {
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)
}
