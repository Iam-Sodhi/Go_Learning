package main

import "fmt"

// execute at last of function, and it  follows LIFO
func main() {
	defer fmt.Println("World")
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("Hello , Defer")
	defer fmt.Println("3")

}
