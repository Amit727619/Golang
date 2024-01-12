package main

import (
	"fmt"
	"strings"
)

func wordcount(text string) map[string]int {
	split := strings.Fields(text) 
	freq := make(map[string]int)

	for _, word := range split {
		word = strings.ToLower(word)
		freq[word]++
	}
	return freq
}

func main() {
	str := "good afternoon everyone , how are you "
	wordfreq := wordcount(str)
	fmt.Println("Word Frequency count ")
	for word, count := range wordfreq{
		fmt.Printf("%s: %d\n", word, count)
	}
}
