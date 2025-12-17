package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"-"` // will always be omitted
}

func main() {
	person := Person{FirstName: "John", LastName: "Doe", Age: 42}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling json:", err)
		return
	}
	fmt.Println(string(jsonData))

}
