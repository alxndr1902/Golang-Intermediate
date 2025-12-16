package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//fmt.Println(rand.Intn(5) + 1)
	for {
		fmt.Println("Welcome to the dice game")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Println("Enter your choice")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, please enter 1 or 2")
			continue
		}
		if choice == 2 {
			fmt.Println("Thank you for playing, good bye")
			break
		}

		dice1 := rand.Intn(6) + 1
		dice2 := rand.Intn(6) + 1

		fmt.Printf("You rolled a %d, and %d.\n", dice1, dice2)
		fmt.Println("Total is:", dice1+dice2)

		fmt.Println("Do you want to play again? y/n: ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, Assuming no")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thanks for playing, goodbye.")
			break
		}
	}
}
