package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := Person{
		Name:  "John Doe",
		Email: "example@gmail.com",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	person2 := Person{
		Name:  "Jane Doe",
		Age:   30,
		Email: "jane@example.com",
		Address: Address{
			City:  "New York",
			State: "New York",
		},
	}

	jsonData1, err := json.Marshal(person2)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData1))

	jsonData2 := `{"full_name": "Jane doe", "emp_id": "001", "age": 30, "address": {"city": "san jose", "state": "CA"}}`

	var employeeFromJSON Employee
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJSON)
	if err != nil {
		fmt.Println("Error unmarshalling to JSON:", err)
		return
	}
	fmt.Println(employeeFromJSON)
	fmt.Println(employeeFromJSON.Age + 5)
	fmt.Println(employeeFromJSON.Address.City)

	// handling unknown json structure
	jsonData3 := `{"name": "John", "age": 30, "address": {"city": "new york", "state": "NY"}}`

	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		log.Fatalln("Error unmarshalling to JSON:", err)
	}
	fmt.Println("Decoded/unmarhsalled JSON", data)
	fmt.Println("Decoded/unmarhsalled JSON", data["address"])
	fmt.Println("Decoded/unmarhsalled JSON", data["name"])
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
