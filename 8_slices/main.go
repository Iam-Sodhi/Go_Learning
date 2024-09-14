package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Banana", "Peach"}

	fmt.Printf("Type is %T \n", fruitList)

	fruitList = append(fruitList, "Mango")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	//3 is not inclusive
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	//Slices:  by default , 4 memory is given
	highScores[0] = 123
	highScores[1] = 234
	highScores[2] = 133
	highScores[3] = 438

	fmt.Println(highScores)

	highScores = append(highScores, 343, 234, 392)
	//would not give error
	fmt.Println(highScores)
	sort.Ints(highScores)
	//to sort in slice
	fmt.Println(highScores)

	var courses = []string{"reactjs", "javascript", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	//use append to remove the elements
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
