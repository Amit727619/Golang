/*
Even or Odd:

Write a program that prompts the user to enter an integer and determines whether the entered number is even or odd. Display an appropriate message based on the result.
*/
package main

import (
	"fmt"
)

func isEven(number int) bool {
	return number%2 == 0
}

func main() {
	var num int

	fmt.Print("Enter an integer: ")
	fmt.Scan(&num)

	if isEven(num) {
		fmt.Printf("%d is an even number!\n", num)
	} else {
		fmt.Printf("%d is an odd number.\n", num)
	}
}
