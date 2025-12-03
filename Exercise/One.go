package main

import "fmt"

func main() {
	var X int
	var Y int
	var Z float64
	var result float64
	fmt.Println("System akan menghitung ((X+Y+Z) * (Z/2))-Z")
	fmt.Println("Input value X dalam bilangan bulat")
	fmt.Scanln(&X)
	fmt.Println("Input value Y dalam bilangan bulat")
	fmt.Scanln(&Y)
	fmt.Println("Input value Z dalam bilangan decimal")
	fmt.Scanln(&Z)
	fmt.Println("Hasilnya adalah")
	result = ((float64(X) + float64(Y) + Z) * (Z / 2)) - Z
	fmt.Println(result)

	fmt.Println("Ubah value X menjadi:")
	fmt.Scanln(&X)
	fmt.Println("Hasilnya sekarang adalah")
	result = ((float64(X) + float64(Y) + Z) * (Z / 2)) - Z
	fmt.Println(result)

	fmt.Println("Ubah value Y menjadi:")
	fmt.Scanln(&Y)
	fmt.Println("Hasilnya sekarang adalah")
	result = ((float64(X) + float64(Y) + Z) * (Z / 2)) - Z
	fmt.Println(result)

	fmt.Println("Ubah value Z menjadi:")
	fmt.Scanln(&Z)
	fmt.Println("Hasilnya sekarang adalah")
	result = ((float64(X) + float64(Y) + Z) * (Z / 2)) - Z
	fmt.Println(result)
}
