package main

import (
	"fmt"
	"encoding/json"
)
//SensorReading ...
type SensorReading struct { 
	Name 		string 	`json:"name"`
	Capacity 	int 	`json:"capacity"`
	Time 		string	`json:"time"`
	Information Info 	`json:"info"`
}

// Info ...
type Info struct {
	Description string `json:"desc"`
}

func main() {
	fmt.Println("Hello World")	

	jsonSting := `{
		"name": "battery sensor", 
		"capacity": 40,
		"time": "2019-01-21T19:07:28Z",
		"info": {
			"desc": "a sensor reading"
		}
	}`

	var reading SensorReading
	err := json.Unmarshal([]byte(jsonSting), &reading)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)
}





// convert data to JSON
// package main

// import (
// 	"fmt"
// 	"encoding/json"
// )

// type Book struct {
// 	Title  string `json:"title"`
// 	Author Author `json:"author"`
// }

// type Author struct {
// 	Name	 	string `json:"name"`
// 	Age 		int `json:"age"`
// 	Developer 	bool `json:"is_developer"`
// }

// func main() {
// 	fmt.Println("Hello World")

// 	author:= Author{Name: "Elliot", Age: 25, Developer: true}
// 	book := Book{Title: "Learning Concurrency in Python", Author: author}

// 	fmt.Printf("%+v\n", book)

// 	byteArray, err := json.MarshalIndent(book,"","	")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(string(byteArray))
// }
