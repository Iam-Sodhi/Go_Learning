package main

import "fmt"

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// fmt.Println("using range in for")
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 5 {
		if rougueValue == 3 {
			goto leo
		}
		fmt.Printf("Value is: %v\n", rougueValue)
		rougueValue++

	}

leo:
	fmt.Println("Jumping to leo sign")

}
