package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	//checkError(os.Mkdir("subdir", 0755))
	////defer os.RemoveAll("subdir")
	//os.WriteFile("subdir/file", []byte(""), 0755)
	//
	//checkError(os.MkdirAll("subdir/parent/child", 0755))
	//checkError(os.MkdirAll("subdir/parent/child1", 0755))
	//checkError(os.MkdirAll("subdir/parent/child2", 0755))
	//checkError(os.MkdirAll("subdir/parent/child3", 0755))
	//os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}
}
