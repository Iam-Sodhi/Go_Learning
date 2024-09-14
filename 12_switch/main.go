package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dive value is 1")
	case 2:
		fmt.Println("DV=2")

	case 3:
		fmt.Println("Dive value is 3")
	case 4:
		fmt.Println("DV=4")
	case 5:
		fmt.Println("Dive value is 5")
	case 6:
		fmt.Println("DV=6")
	default:
		fmt.Println("What was that")

	}
}
