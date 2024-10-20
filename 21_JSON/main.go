package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //- will remove the field from json
	Tags     []string `json:"tags,omitempty"` //if value nil/null, omitempty does not show that field
}

func main() {
	fmt.Println("Welcome to JSON...")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	s_courses := []course{
		{"Reactjs", 111, "http.in.react", "admin", []string{"web", "js"}},
		{"AI/ML learn", 133, "http.in.ml", "adminq23", nil},
	}

	//package this data as JSON data
	// finalJson, err := json.Marshal(s_courses)
	finalJson, err := json.MarshalIndent(s_courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {

	json_data := []byte(`
	  {
	 "coursename": "reactjs",
	  "Price": 299,
	  "website":  "http.in.react",
      "tags": ["web","js"]
	  }
	`)

	var s_courses course

	checkValid := json.Valid(json_data) //to check if data coming is valid

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(json_data, &s_courses)
		fmt.Printf("%#v\n", s_courses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//some cases where you just want to add data to key value
	var my_data map[string]interface{}
	json.Unmarshal(json_data, &my_data)
	fmt.Printf("%#v\n", my_data)

	for k, v := range my_data {
		fmt.Printf("Key is %v and value is %v\n", k, v)
	}
}
