package main


import ("fmt")

type Person struct{
	name : string
	addr : string
	age : int
}

func main(){
	var pers1 Person
	var pers2 Person


	pers1.name = "durvesh"
	pers1.addr = "pune"
	pers1.age = 27

	fmt.Println("name",pers1.name)
	fmt.Println("addr",pers1.addr)
	fmt.Println("age",pers1.age)

}