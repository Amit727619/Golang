
package main 

import "fmt"

type Car struct {
	Name , Model , Color string 
    Cost int
}

func main(){
	c := Car{Name: "skoda",Model: "skoda",Color "red",cost "18,00000"}
    
	fmt.Prinln("car name", c.Name)
    fmt.Prinln("car color",c.Color)

	c.Color = "Black"

	fmt.Println("Car new color", c)


}