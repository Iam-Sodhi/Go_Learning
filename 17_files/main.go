package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	content := "This needs to go in a file"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err) //panic will shut down execution of program and shows err
	}
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close() //recommended to use defer

	readFile("./myfile.txt")
}

func readFile(filepath string) {
	databyte, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is \n", string(databyte))
}
