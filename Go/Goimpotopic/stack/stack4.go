package main

import "fmt"

func main() {
    stack := []int{1, 2, 3, 4, 5}

    // Create a new slice excluding the last element
    newSlice := stack[:len(stack)-1]

    fmt.Println("Original Stack:", stack)
    fmt.Println("New Slice (excluding the last element):", newSlice)
}