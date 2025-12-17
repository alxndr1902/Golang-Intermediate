package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed subdir
var subdir embed.FS

func main() {
	//fmt.Println("Embedded content:", content)
	content, err := subdir.ReadFile("subdir/example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Embedded file content:\n", string(content))

	err = fs.WalkDir(subdir, "subdir", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("Path:", path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
