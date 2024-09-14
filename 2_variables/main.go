package main

import "fmt"

// jwtT := 23333   //here we can't use := outside any method only = is allowed here

func main() {

	var username string = "gautam"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 // unsigned, we can't use more than 255 value
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.23432423423423423423423234 // less precision than float64
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar) //garbage value will always be zero when you only declare
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	//implicit type of declaring variable
	var website = "gautam-sodhi.vercel.app"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website) //here lexer will automatically define its type and further we would n't be able to change its type

	// no var style
	numberofUser := 303200.0 //it is a short variable declaration(not assignement like =) & we must declare at least one new variable when using it
	fmt.Println(numberofUser)

}
