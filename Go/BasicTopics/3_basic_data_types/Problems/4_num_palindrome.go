/*
Number Palindrome:

Create a program that checks whether a given integer is a palindrome. A palindrome is a number that reads the same backward as forward. For example, 121 is a palindrome.
*/

package main

import (
	"fmt"
)

func isPalindrome(number int) bool {
	original := number
	reversed := 0

	for number > 0 {
		digit := number % 10
		reversed = reversed*10 + digit
		number /= 10
	}

	return original == reversed
}

func main() {
	var num int

	fmt.Print("Enter an integer: ")
	fmt.Scan(&num)

	if isPalindrome(num) {
		fmt.Printf("%d is a palindrome!\n", num)
	} else {
		fmt.Printf("%d is not a palindrome.\n", num)
	}
}
