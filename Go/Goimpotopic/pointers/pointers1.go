package main
import "fmt"
func main() {
    var num = 20 
    var intPointer *int
  
    intPointer = &num

    fmt.Println(" num ", num)
    fmt.Println("address of num", &num)
    fmt.Println(" intPointer stores:", intPointer)
}