// The Spiral Matrix problem involves traversing a 2D matrix in a spiral order, starting from the top-left corner and moving in a clockwise direction. The task is to return all elements of the matrix in the order of their traversal. Here's a more detailed explanation:

// # Problem Statement:
// Given an m x n matrix, return all elements of the matrix in spiral order.

// # Example:
// Consider the following 3x3 matrix:

// 1 2 3
// 4 5 6
// 7 8 9

// The spiral order traversal of this matrix would be [1, 2, 3, 6, 9, 8, 7, 4, 5].


// feedback

// Correctness: 5/5
// The program correctly generates the spiral order of the input matrix. It follows the correct order of traversing the matrix in a spiral manner, covering all elements exactly once.

// Complexity: 4/5
// The time complexity of the program is O(m * n), where m is the number of rows and n is the number of columns in the matrix. This is because the program traverses each element once. The space complexity is O(1) as the program uses a constant amount of extra space regardless of the input size. The additional space used for the output array is not considered in the space complexity analysis.

// Test Cases: 4/5
// The provided test case covers a basic scenario with a 3x3 matrix. It would be beneficial to have additional test cases to validate the program's correctness for edge cases, such as a 1x1 matrix, a matrix with a single row or single column, an empty matrix, and larger matrices.

// Workflow: 5/5
// The program follows a clear and well-structured workflow. The logic is organized, making it easy to understand. The use of meaningful variable names enhances readability. The loop structure appropriately handles the traversal in a spiral order.

// Overall, the program is correct and well-structured. Consider adding more test cases to ensure the program's correctness for various scenarios, and the space complexity could be further clarified by accounting for the additional space used by the output array.


package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	output := spiralOrder(matrix)
	fmt.Println(output)

}

func spiralOrder(matrix [][]int) []int {

	var output []int
	totalRows := len(matrix)
	totalColumns := len(matrix[0])

	startRow := 0
	endRow := totalRows - 1
	startColumn := 0
	endColumn := totalColumns - 1

	for startRow <= endRow && startColumn <= endColumn {
          
		for i := startColumn; i <= endColumn; i++ {
			                           
			output = append(output, matrix[startRow][i])
			                                
		}                                                  
		startRow++  

		for i := startRow; i <= endRow; i++ {
			             
			output = append(output, matrix[i][endColumn])
			                              
		}

		endColumn--
	     	

		if startRow <= endRow {
			
			for i := endColumn; i >= startColumn; i-- {   
                       
				output = append(output, matrix[endRow][i])
			}                                     
		}

		endRow--  

		if startColumn <= endColumn {
			
			for i := endRow; i >= startRow; i-- {
                                
				output = append(output, matrix[i][startColumn])
			}
		}
		startColumn++
	}

	return output

}