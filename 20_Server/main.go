package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8080/postform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "gautam")
	data.Add("lastname", "sodhi")
	data.Add("email", "a@b.com")

	response, err := http.PostForm(myUrl, data) //here we don't need to tell type of data

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
func PerformPostRequest() {
	const myUrl = "http://localhost:8080/post"

	//fake json payload
	requestBody := strings.NewReader(`
	  {
		"courseName" : "Golang",
		"price": 0,
		"platform": "youTube"
	   }
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8080/"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	// content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(content)
	// fmt.Println(string(content))

	//another method
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

}
