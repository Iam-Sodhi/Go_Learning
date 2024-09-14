package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"] = "Web dev"
	languages["Py"] = "Python"
	languages["c++"] = "Low level language"

	fmt.Println(languages)
	fmt.Println("C++ shorts for: ", languages["c++"])

	//to delete
	delete(languages, "Py")
	fmt.Println(languages)
}
