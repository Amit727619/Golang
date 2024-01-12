package main

import (
	"fmt"
)

func countCharacters(str string) map[rune]int {
	charCount := make(map[rune]int)

	for _, char := range str {
		charCount[char]++
	}

	return charCount
}

func main() {
	String := "Hello, World!"
	a := countCharacters(String)

	fmt.Println("Character occurrences:")
	for char, count := range a {
		fmt.Printf("character %c count  %d\n", char, count)
	}
}
