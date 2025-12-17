package main

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//tempFile, err := os.CreateTemp("", "temporaryFile")
	//CheckError(err)
	//fmt.Println("Temporary file created:", tempFile.Name())
	//
	//defer tempFile.Close()
	//defer os.Remove(tempFile.Name())

	tempDir, err := os.MkdirTemp("", "temporaryDir")
	CheckError(err)
	fmt.Println("Temporary directory created:", tempDir)
	defer os.Remove(tempDir)
}
