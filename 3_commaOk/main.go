package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter some number: ")

	//comman ok  OR Error Ok syntax

	// input, err=
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for answering, ", input)

}
