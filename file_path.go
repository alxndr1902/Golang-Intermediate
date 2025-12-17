package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// join path using filepath.Join()
	joinedPath := filepath.Join("Documents", "Download", "file.zip")
	fmt.Println("Joined path:", joinedPath)

	// clean path using filepath.Clean()
	normalizedPath := filepath.Clean("./Data/../data/file.txt")
	fmt.Println("Normalized path:", normalizedPath)

	dir, file := filepath.Split("/home/users/docs/file.txt")
	fmt.Println("File:", file)
	fmt.Println("Path:", dir)
	fmt.Println(filepath.Base("/home/users/docs/file.txt"))

	relativePath := "./data/file.txt"
	fmt.Println(filepath.IsAbs(relativePath))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Absolute path:", absPath)
	}
}
