package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println(ptr)    //will print <nil>

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("The address is ", ptr)
	fmt.Println("The value is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("The new value is", *ptr)

}
