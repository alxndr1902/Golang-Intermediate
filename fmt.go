package main

import "fmt"

func main() {
	fmt.Print("hello\n")
	fmt.Print("world!\n")

	var name string
	var age int

	fmt.Print("Enter your name and age:")
	//fmt.Scan(&name, &age)
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("name: %s, age: %d\n", name, age)

	//error formatting functions
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age: %d is too young to drive\n", age)
	}
	return nil
}
