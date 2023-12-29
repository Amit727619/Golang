/*
Problem:
Temperature Conversion:
Write a program that converts a temperature in Celsius to Fahrenheit. Prompt the user to enter a temperature in Celsius, perform the conversion, and print the result in Fahrenheit.
*/

package main

import (
	"fmt"
)

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9/5) + 32
}

func main() {
	var celsius float64

	fmt.Print("Enter temperature in Celsius: ")
	fmt.Scan(&celsius)
	fahrenheit := celsiusToFahrenheit(celsius)
	
	fmt.Printf("%.2f degrees Celsius is equal to %.2f degrees Fahrenheit\n", celsius, fahrenheit)
}
