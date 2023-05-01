/*
1491. Average Salary Excluding the Minimum and Maximum Salary
You are given an array of unique integers salary where salary[i] is the salary of the ith employee.

Return the average salary of employees excluding the minimum and maximum salary. Answers within 10-5 of the actual answer will be accepted.

Example 1:

Input: salary = [4000,3000,1000,2000]
Output: 2500.00000
Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
Average salary excluding minimum and maximum salary is (2000+3000) / 2 = 2500
Example 2:

Input: salary = [1000,2000,3000]
Output: 2000.00000
Explanation: Minimum salary and maximum salary are 1000 and 3000 respectively.
Average salary excluding minimum and maximum salary is (2000) / 1 = 2000
*/
package main

import "fmt"

func average(salary []int) float64 {
	if len(salary) <= 2 {
		return 0
	}

	min, max := salary[0], salary[1]

	if min > max {
		min, max = max, min
	}

	res := 0.0
	for i := 0; i < len(salary); i++ {
		if min > salary[i] {
			min = salary[i]
		}

		if max < salary[i] {
			max = salary[i]
		}
		res += float64(salary[i])
	}

	res -= float64(min)
	res -= float64(max)

	return res / float64(len(salary)-2)
}

func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
}
