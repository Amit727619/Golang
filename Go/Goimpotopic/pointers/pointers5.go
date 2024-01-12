
package main

import "fmt"

func display() *string{

	m := "morning"

	return &m
}


func main(){
	
	str := display()
	fmt.Prinln("good",*str)

}