/*
54. Spiral Matrix
Given an m x n matrix, return all elements of the matrix in spiral order.

Example 1:
Input: matrix = [[1,2,3],
                [4,5,6],
                [7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

package main

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}

	rows, cols := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, cols-1, 0, rows-1

	for left <= right && top <= bottom {
		// traverse left
		for col := left; col <= right; col++ {
			res = append(res, matrix[top][col])
		}
		// traverse dow
		for row := top + 1; row <= bottom; row++ {
			res = append(res, matrix[row][right])
		}

		// Traverse left if required
		if left < right && top < bottom {
			for col := right - 1; col > left; col-- {
				res = append(res, matrix[bottom][col])
			}
			// Traverse up
			for row := bottom; row > top; row-- {
				res = append(res, matrix[row][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return res

}
