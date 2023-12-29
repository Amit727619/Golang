/*
BMI Calculator:

Write a program that calculates the Body Mass Index (BMI) based on the user's weight in kilograms and height in meters. Prompt the user to enter their weight and height, perform the calculation, and display the BMI.

BMI = weight(kg)/height(m) ** 2

*/

package main

import (
	"fmt"
)

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func main() {
	var weight, height float64
	fmt.Print("Enter your weight in kilograms: ")
	fmt.Scan(&weight)

	fmt.Print("Enter your height in meters: ")
	fmt.Scan(&height)

	bmi := calculateBMI(weight, height)

	fmt.Printf("Your BMI is: %.2f\n", bmi)
}
