//  Problem Statement:
// Given an array of integers, find the contiguous subarray (containing at least one number) with the largest sum and return the sum.

// Example:
// Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]

// Output: subarray [4, -1, 2, 1] maxSum 6

// Explanation: The contiguous subarray [4, -1, 2, 1] has the largest sum of 6. 


// feedback

// Correctness: 5/5
// The program correctly identifies the contiguous subarray with the largest sum using the Kadane's algorithm. The logic for finding the subarray and calculating the maximum sum seems correct.

// Complexity: 4/5
// The time complexity of the program is O(n), where n is the length of the input array. This is due to a single pass through the array. The space complexity is O(1), as the program uses a constant amount of extra space regardless of the input size. However, the calculation of the subarray using slicing may introduce an additional O(n) space complexity, impacting space efficiency.

// Test Cases: 4/5
// The provided test case covers a basic scenario with a mix of positive and negative numbers. However, it would be beneficial to have additional test cases that include edge cases, such as an array with all negative numbers, all positive numbers, a single-element array, and an empty array. More extensive testing would enhance the robustness of the program.

// Workflow: 5/5
// The program follows a clear and straightforward workflow. The logic is well-structured, making it easy to understand. The use of meaningful variable names and comments further enhances the readability of the code.

// Overall, the program is well-implemented, but there is room for improvement in terms of additional test cases for thorough testing and addressing the space complexity concern related to subarray slicing.
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
