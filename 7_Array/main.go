package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array part")

	var fruitList [4]string
	fruitList[0] = "Apple0"
	fruitList[1] = "Apple1"
	fruitList[3] = "Apple3"

	fmt.Println("FruitList is: ", fruitList)
	fmt.Println("FruitList is: ", len(fruitList))

	var vegList = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println("VegList is: ", vegList)
}
