package main
import ("fmt")

type Person struct{
	var name string 
	var addr string 
	var age int 
}
func main(){
	var pers1 Person
	var pers2 Person

	pers1.name = "durevsh"
    pers1.addr = "pune"
	pers1.age = 27

   printPerson(pers1)

   printPerson(pers2)
   
}

func printPerson(pers Person){
	fmt.Print("name",pers.name)
	fmt.Print("addr",perse.addr)
	fmt.Printf("age ",perse.age)
}