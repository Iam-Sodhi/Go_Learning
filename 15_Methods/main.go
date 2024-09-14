package main

import "fmt"

func main() {
	//no inheritance in golang;
	// No super or parent concepts
	gautam := User{"Gautam", "gautam@go.dev", true, 16}
	fmt.Println(gautam)
	fmt.Printf("gautam details are: %+v\n", gautam)
	gautam.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() { // here only copy is being passed of user.so change here would not reflect on original unless we use pointers
	fmt.Println("is user active: ", u.Status)
}
