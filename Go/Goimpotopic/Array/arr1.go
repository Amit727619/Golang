 
package main 

import "fmt"

func main() { 
	 
arr1:= [4]int{1,2,3,4} 
arr2:= [...]int{7,3,2,4,8,9} 
arr3:= [7]int{1,4,5,2,5,6,8} 

fmt.Println("length arr1", len(arr1)) 
fmt.Println("length arr2", len(arr2)) 
fmt.Println("length arr3", len(arr3)) 
} 
