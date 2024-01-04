//  Problem Statement:
// Given an array of integers, find the contiguous subarray (containing at least one number) with the largest sum and return the sum.

// Example:
// Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]

// Output: subarray [4, -1, 2, 1] maxSum 6

// Explanation: The contiguous subarray [4, -1, 2, 1] has the largest sum of 6. 

package main

import (
	"fmt"
)

func maxSubArray(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, nil
	}

	maxSum := nums[0]
	currentSum := nums[0]
	startIndex, tempStart := 0, 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > currentSum+nums[i] {
			currentSum = nums[i]
			tempStart = i
		} else {
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
			startIndex = tempStart
		}
	}

	// Finding the subarray
	subarray := nums[startIndex : startIndex+len(nums)]

	return maxSum, subarray
}

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum, subarray := maxSubArray(input)

	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Output: subarray %v maxSum %v\n", subarray, maxSum)
}
