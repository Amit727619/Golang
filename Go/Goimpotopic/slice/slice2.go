package main

import ("fmt")

func main() {
	myslice1 := []int{1,2,3,4}
	myslice2 := []int{5,6,7,8}
	myslice3  := append(myslice1,myslice2...)

	fmt.Println("myslice3 %v \n",myslice3)
	fmt.Println("length  %d \n",len (myslice3))
	fmt.Println("capacity %d \n",cap (myslice3))

}