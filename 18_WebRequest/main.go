package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://gautam-sodhi.vercel.app"

func main() {
	fmt.Println("Web Request: ")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() //caller's responsibilty to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("The respnse is : ", content)

}
