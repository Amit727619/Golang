package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	fmt.Println("mid",mid)
	left := mergeSort(arr[:mid])
	fmt.Println("left",left)
	right := mergeSort(arr[mid:])
	fmt.Println("right",right)
	fmt.Println("merge sort function")
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			fmt.Println(result)
			i++
		} else {
			result = append(result, right[j])
			fmt.Println(result)
			j++
		}
	}

	result = append(result, left[i:]...)
	fmt.Println(result)
	result = append(result, right[j:]...)
	fmt.Println(result)

	return result
}

func main() {
	arr := []int{4,2,8,6}
	fmt.Println("input array", arr)

	sort := mergeSort(arr)
	fmt.Println("array aftr sorting ", sort)
}
