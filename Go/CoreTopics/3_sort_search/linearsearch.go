package main

import "fmt"

func linearSearch(arr []int, target int) int {
    for i := 0; i < len(arr); i++ {
        if arr[i] == target {
            return i 
        }
    }
    return -1 
}

func main() {
   
    arr := []int{1, 2, 3, 4}
    target := 4
    index := linearSearch(arr, target)

    if index != -1 {
        fmt.Printf("target %d found at index %d\n", target, index)
    } else {
        fmt.Printf("target %d not found \n", target)
    }
}
