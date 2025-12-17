package main

import (
	"flag"
	"fmt"
)

func main() {
	//fmt.Println("Command:", os.Args[0])
	//
	//for _, arg := range os.Args {
	//	fmt.Println("Arguments:", arg)
	//}

	// define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "John Doe", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of the user")
	flag.BoolVar(&male, "male", true, "Gender of the user")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)
}
