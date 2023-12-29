/*
Write a program that reads a text file and counts the number of lines in it. Use bufio.Reader for efficient line-by-line reading.
*/



package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	Count := 1

	for {
		_, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return 0, err
		}

		Count++
	}

	return Count, nil
}

func main() {
	filename := "output.txt"

	Count, err := countLines(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Number of lines in the file: %d\n", Count)
}
