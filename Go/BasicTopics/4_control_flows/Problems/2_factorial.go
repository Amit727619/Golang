/* Factorial Calculation:

Write a program that calculates the factorial of a given non-negative integer. The factorial of a number is the product of all positive integers up to that number. For example, the factorial of 5 is 5 * 4 * 3 * 2 * 1.
*/


package main

import (
	"fmt"
)

func calculateFactorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * calculateFactorial(n-1)
}

func main() {
	var num int

	fmt.Print("Enter a integer: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Please enter a non-negative integer.")
		return
	}

	factorial := calculateFactorial(num)
	
	fmt.Printf("The factorial of %d is: %d\n", num, factorial)
}
