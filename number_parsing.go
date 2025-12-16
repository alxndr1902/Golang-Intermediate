package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed integer:", num)
	fmt.Println("Parsed integer:", num+1)

	numistr, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed integer:", numistr)

	floatStr := "3.14"
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}

	fmt.Printf("Parsed float: %.1f\n", floatVal)
}
