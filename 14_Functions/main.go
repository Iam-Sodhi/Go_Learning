package main

import "fmt"

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi pro Result function"
}

func main() {
	fmt.Println("Welcome to Functions in Go")

	result := adder(4, 5)
	fmt.Println("Result of adding 4 & 5: ", result)

	proRes, _ := proAdder(2, 3, 4, 1)
	fmt.Println("The proRes is: ", proRes)
}
