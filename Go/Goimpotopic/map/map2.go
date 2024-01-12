
package main 

import ( 
	"fmt"; 
	"strings"
) 

func main() { 
	modified := func(r rune) rune { 
		if r == 'o' { 
			return '@'
		} 
		return r 
	} 

	str := "good morning"
	fmt.Println(str) 
	result := strings.Map(modified, str) 
	fmt.Println(result)
	spaceremove()
} 

func spaceremove() { 
	transformed := func(r rune) rune { 
		
		if r == ' ' { 
			return 0 
		} 
		return r 
	} 

	str := "good morning how are you"
	fmt.Println(str) 

	output := strings.Map(transformed, str) 
	fmt.Println(output) 
} 


