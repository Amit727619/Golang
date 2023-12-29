/*
Count Vowels and Consonants:

Create a function that takes a string as input and returns the count of vowels and consonants in the string. Ignore spaces and consider case-insensitivity.
*/

package main

import "fmt"
func main() {
   
   var sentence string

   var vowels, consonant int

   sentence = "Good Morning"
   
   vowels = 0

   consonant = 0
   
   for i := 0; i < len(sentence); i++ {
      if sentence[i] == ' ' {
         continue
      }
   
     
      if sentence[i] == 'a' || sentence[i] == 'e' || sentence[i] == 'i' || sentence[i] == 'o' || sentence[i] == 'u' ||
         sentence[i] == 'A' || sentence[i] == 'E' || sentence[i] == 'I' || sentence[i] == 'O' || sentence[i] == 'U' {

            vowels++
      } else {
       
         consonant++
      }
   }
   fmt.Println("Sentence:- \n", sentence)
   fmt.Println("Result:- \nThe total number of vowels ", vowels)
   fmt.Println("The total number of consonants ", consonant)
}
