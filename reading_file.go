package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer func() {
		fmt.Println("Closing opened file")
		err := file.Close()
		if err != nil {
			return
		}
	}()

	fmt.Println("File opened successfully")

	//// read the contents of the opened file
	//data := make([]byte, 1024)
	//_, err = file.Read(data)
	//if err != nil {
	//	fmt.Println("error reading data from file:", err)
	//	return
	//}
	//
	//fmt.Println("File content:", string(data))

	scanner := bufio.NewScanner(file)

	// Read line by line
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d: %s\n", lineNumber, line)
		lineNumber++
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
