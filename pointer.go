package main

import "fmt"

func main() {
	// Pointer is a variable that stores the memory address of another variable
	var ptr *int
	var a int = 10
	ptr = &a //referencing
	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr) //dereferencing
	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++
}
