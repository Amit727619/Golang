
package main

import "fmt"

// defining a structure
type Employee struct {
	firstName, lastName string
	age, salary int
}

func main() {

	// passing the address of struct variable
	// emp8 is a pointer to the Employee struct
	emp8 := &Employee{"Sam", "foundry", 55, 6000}

	fmt.Println("first name :", (*emp8).firstName)
	fmt.Println("age:", (*emp8).age)
}
