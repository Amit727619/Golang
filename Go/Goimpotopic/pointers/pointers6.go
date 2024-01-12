
package main
import "fmt"


func callByvalue(num int){
	num =30
	fmt.Println(num)
}
func callByreference(num *int){
	*num = 20
	fmt.Prinln(*num)
}
func main(){

	var number int 

	callByvalue(number)
	callByreference(&number)
}