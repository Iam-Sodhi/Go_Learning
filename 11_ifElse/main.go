package main

import "fmt"

func main() {
	loginCount := 10

	var result string
	// {   needs to be here . not on second line.
	if loginCount < 10 {
		result = "regular useer"
	} else if loginCount > 10 {
		result = "WatchOut"
	} else {
		result = "Consistent"
	}

	fmt.Println(result)

	if num := 3; num < 4 {
		fmt.Println("less than 4")
	} else {
		fmt.Println("No. >=4")
	}
}
