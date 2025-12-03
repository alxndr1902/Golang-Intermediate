package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(sumOfDigits(11))
}
func factorial(n int) int {
	// base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// recursive case: factorial of n is n * factorial(n-1)
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}
