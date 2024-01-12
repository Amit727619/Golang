package main
import "fmt"

func main() {
  var num  int
  var ptr *int
    
  num = 2
  fmt.Println("Address of num:",&num)
  fmt.Println("Value of num:",num)

  ptr = &num
  fmt.Println("Address of pointer ptr:",ptr)
  fmt.Println("Content of pointer ptr:",*ptr)
    
  num = 1
  fmt.Println(" Address of pointer ptr:",ptr)
  fmt.Println("Content of pointer ptr:",*ptr)
    
  *ptr = 4
  fmt.Println(" Address of num:",&num)
  fmt.Println("Value of num:",num)
}