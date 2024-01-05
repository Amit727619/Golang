// The "Find All Anagrams in a String" problem involves searching for all occurrences of anagrams of a given non-empty string p within a larger string s. An anagram is a word or phrase formed by rearranging the letters of another, such as "cinema" and "iceman." In this problem, you are required to find all starting indices of anagrams of p in s.

// Problem Statement:
// Given a string s and a non-empty string p, find all the starting indices of anagrams of p in s.

// Example:
// plaintext
// Copy code
// Input: s = "cbaebabacd", p = "abc"
// Output: [0, 6]
// Explanation:
// The substring with starting index 0 is "cba," which is an anagram of "abc."
// The substring with starting index 6 is "bac," which is an anagram of "abc

package main

import "fmt"


func Anagrams(str, p string) bool {
    if len(str) != len(p) {
        return false
    }

    charCount := make(map[rune]int)

   
    for i := 0; i < len(str); i++ {
		char := rune(str[i])
		charCount[char]++
	}

    for i := 0; i < len(p); i++ {
		char := rune(p[i])
		charCount[char]--
	}


       
    for i := 0; i < len(charCount); i++ {
        if charCount[rune(i)] != 0 {
        return false
        }
    }


    return true
}
  

func main() {
    str := "cbaebabacd"
     p :="abc"

    for i := 0; i < len(str)-2; i ++ {
        str1 := str[i : i+3]
        
        fmt.Printf("%s ", str1)
        result := Anagrams(str, p)
        fmt.Println(result)
    }
    
}
