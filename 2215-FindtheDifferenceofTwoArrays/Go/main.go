/*
2215. Find the Difference of Two Arrays

Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.

Example 1:
Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Explanation:
For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].
Example 2:

Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
Output: [[3],[]]
Explanation:
For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
Every integer in nums2 is present in nums1. Therefore, answer[1] = [].
*/
package main

func generateMap(nums []int) map[int]bool {
	mp := make(map[int]bool)

	for _, num := range nums {
		mp[num] = true
	}
	return mp
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	mpNums1, mpNums2 := generateMap(nums1), generateMap(nums2)

	var res1 []int
	for _, num := range nums1 {
		if _, ok := mpNums2[num]; !ok {
			mpNums2[num] = true
			res1 = append(res1, num)
		}
	}
	var res2 []int
	for _, num := range nums2 {
		if _, ok := mpNums1[num]; !ok {
			mpNums1[num] = true
			res2 = append(res2, num)
		}
	}

	res := make([][]int, 2)
	res[0] = res1
	res[1] = res2
	return res
}
