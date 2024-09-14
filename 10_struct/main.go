package main

import "fmt"

func main() {
	//no inheritance in golang;
	// No super or parent concepts
	gautam := User{"Gautam", "gautam@go.dev", true, 16}
	fmt.Println(gautam)
	fmt.Printf("gautam details are: %+v\n", gautam)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
