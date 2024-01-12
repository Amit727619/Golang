// Golang Pointer to an Array as Function Argument


package main 

import "fmt"

func updatearray(funarr *[5]int) { 
	(*funarr)[4] = 2
} 

func main() { 

	arr := [5]int{1, 2, 3, 4, 5} 
	updatearray(&arr) 
	fmt.Println(arr) 
} 
