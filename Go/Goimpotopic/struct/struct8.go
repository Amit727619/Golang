
package main 

import ( 
	"fmt"
	"sort"
) 

func main() { 

	// a structure 
	Student := []struct { 
		a_name string 
		a_age int
		a_id  int
	}{ 
		
		{"Durvesh", 26, 1}, 
		{"Mohit", 28, 2}, 
		{"Shubahm", 19, 3}, 
		{"Kritika", 20, 4}, 
	} 

	sort.Slice(Student, func(p, q int) bool { 
	return Student[p].a_name < Student[q].a_name }) 
	
	fmt.Println("Sort Student by names") 
	fmt.Println(Student) 

	sort.Slice(Student, func(p, q int) bool { 
	return Student[p].a_age < Student[q].a_age }) 
	
	fmt.Println() 
	fmt.Println("Sort Student  according to their"+ 
					" total number of age:") 
	fmt.Println(Student) 

} 
