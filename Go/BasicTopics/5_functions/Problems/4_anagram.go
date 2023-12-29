/*
Check for Anagrams:

Write a function that takes two strings as input and determines whether they are anagrams of each other. Anagrams are words or phrases formed by rearranging the letters of another.
*/

package main
import (
	"fmt"
	"strings"
	
)
func main() {
	var str1, str2 string

	fmt.Print("Enter first string: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter second string: ")
	fmt.Scanln(&str2)


	 var str1 string
     var str2 string
	 if Anagram(str1 ,str2 ){ 
         fmt.println("Anagaram")
	 } else{
		fmt.println("")
	 }

	if Anagrams(str1, str2) {
		fmt.Println("The strings are anagrams!")
	} else {
		fmt.Println("The strings are not anagrams!")
	}
}
func Anagrams(str1, str2 string) bool {
	
	string1 := len(s)
	string2 := len(t)
	fmt.Println("Letter 1 =", s, "\nLetter 2 =", t)
	fmt.Println("Is it an Anagram ?")
	if string1 != string2 {
	   return false
	}
}
	

func Anagram(str1 ,str2 string )bool {
	string 1 :=len(s)
	string 2 :=len(t)
	fmt.Println("le")
	fmt.Println("")
	if string1 != string2 {

	}
func Anagram(str1 ,str2 string )bool{

}
}

anagramMap := make(map[string]int)
	
	for i := 0; i < string1; i++ {
	   anagramMap[string(s[i])]++
	}
	
