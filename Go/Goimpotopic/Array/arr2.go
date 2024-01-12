//Copy an Array into Another Array in Golang?


package main

import "fmt"

func main(){

	arr1 := [5]string{"durvesh","shubham","poonam"}
	arr2 := arr1

	fmt.Println("Array 1"arr1)
	fmt.Println("Array 2"arr2)
    usingloop()	
	copyfunction()
	usingappend()
}




func usingloop(){
	arr1 := []{ 2,3,4,5,8}
	arr2 := make([]int, len(arr1))

	for i value := range arr1{
		arr2 = value
	}
	fmt.println("arr1" arr1)
	fmt.println("arr2" arr2)
}

func copyfunction(){
	arr1 := []{ 2,3,4,5,8}
	arr2 := make([]int, len(arr1))

	copy(arr2,arr1)

	fmt.println("arr1" arr1)
	fmt.println("arr2" arr2)
}

func usingappend(){
	Array := []int{1, 2, 3, 4, 5}
    copyArray := make([]int, 0, len(Array))
 
    copyArray = append(copyArray, Array...)
 
    fmt.Println("Original Array: ", Array)
    fmt.Println("Copy Array: ", copyArray)
}

