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

#include<iostream>
#include<vector>
using namespace std;

vector<int> spiralOrder(vector<vector<int>>& matrix) {
    vector<int> res;
    if(matrix.empty()) return res;

    int rows = matrix.size(), cols = matrix[0].size();
    int left = 0, right = cols -1;
    int top = 0, bottom = rows -1;

    while(left <= right && top <= bottom){
        // Traverse right
        for(int col = left; col<= right; col++){
            res.push_back(matrix[top][col]);
        }

        //Traverse down
        for(int row = top+1; row <= bottom; row++){
            res.push_back(matrix[row][right]);
        }

        // Traverse left if required
        if(left < right && top <bottom){
            for( int col = right-1; col > left; col--){
                res.push_back(matrix[bottom][col]);
            }
            // Traverse up
            for( int row = bottom; row > top; row--){
                res.push_back(matrix[row][left]);
            }
        }
        left++;
        right--;
        top++;
        bottom--;
    }
    return res;
}