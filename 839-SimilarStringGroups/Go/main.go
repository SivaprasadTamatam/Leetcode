/*
Two strings X and Y are similar if we can swap two letters (in different positions) of X, so that it equals Y.
Also two strings X and Y are similar if they are equal.

For example, "tars" and "rats" are similar (swapping at positions 0 and 2), and "rats" and "arts" are similar,
but "star" is not similar to "tars", "rats", or "arts".

Together, these form two connected groups by similarity: {"tars", "rats", "arts"} and {"star"}.
Notice that "tars" and "arts" are in the same group even though they are not similar.
Formally, each group is such that a word is in the group if and only if it is similar to at least one other word in the group.

We are given a list strs of strings where every string in strs is an anagram of every other string in strs. How many groups are there?

Example 1:

Input: strs = ["tars","rats","arts","star"]
Output: 2
Example 2:

Input: strs = ["omv","ovm"]
Output: 1

Constraints:

1 <= strs.length <= 300
1 <= strs[i].length <= 300
strs[i] consists of lowercase letters only.
All words in strs have the same length and are anagrams of each other.
*/
package main

import "fmt"

func CheckIsSimilar(s1, s2 string) bool {
	if s1 == s2 { // two strings X and Y are similar if they are equal.
		return true
	}

	diff := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff == 2 // Two strings X and Y are similar if we can swap two letters (in different positions) of X, so that it equals Y.
}

func dfs(cp int, visited []bool, adjList [][]int) {
	visited[cp] = true
	for _, np := range adjList[cp] {
		if !visited[np] {
			dfs(np, visited, adjList)
		}
	}

}

func numSimilarGroups(strs []string) int {
	ans, n := 0, len(strs)
	adjList := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if CheckIsSimilar(strs[i], strs[j]) {
				adjList[i] = append(adjList[i], j)
				adjList[j] = append(adjList[j], i)
			}
		}
	}

	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			ans++
			dfs(i, visited, adjList)
		}
	}
	return ans
}

func main() {
	strs := []string{"tars", "rats", "arts", "star"}

	strs2 := []string{"omv", "ovm"}
	fmt.Println(numSimilarGroups(strs))
	fmt.Println(numSimilarGroups(strs2))
}
