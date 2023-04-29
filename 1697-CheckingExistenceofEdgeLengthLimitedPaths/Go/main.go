/*
1697. Checking Existence of Edge Length Limited Paths
Hard
1.3K
33
Companies
An undirected graph of n nodes is defined by edgeList, where edgeList[i] = [ui, vi, disi] denotes an edge between nodes ui and vi with distance disi. Note that there may be multiple edges between two nodes.

Given an array queries, where queries[j] = [pj, qj, limitj], your task is to determine for each queries[j] whether there is a path between pj and qj such that each edge on the path has a distance strictly less than limitj .

Return a boolean array answer, where answer.length == queries.length and the jth value of answer is true if there is a path for queries[j] is true, and false otherwise.

Example 1:

Input: n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0,2,5]]
Output: [false,true]
Explanation: The above figure shows the given graph. Note that there are two overlapping edges between 0 and 1 with distances 2 and 16.
For the first query, between 0 and 1 there is no path where each distance is less than 2, thus we return false for this query.
For the second query, there is a path (0 -> 1 -> 2) of two edges with distances less than 5, thus we return true for this query.
Example 2:

Input: n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],[1,4,13]]
Output: [true,false]
Exaplanation: The above figure shows the given graph.

Constraints:

2 <= n <= 105
1 <= edgeList.length, queries.length <= 105
edgeList[i].length == 3
queries[j].length == 3
0 <= ui, vi, pj, qj <= n - 1
ui != vi
pj != qj
1 <= disi, limitj <= 109
There may be multiple edges between two nodes.
*/
package main

import (
	"sort"
)

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	// Add index, so we can fill answer later proper after sorting
	for i := 0; i < len(queries); i++ {
		queries[i] = append(queries[i], i)
	}

	// Sort according to distance
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	// Sort according to distance
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})

	parent := make([]int, n)
	rank := make([]int, n)
	res := make([]bool, len(queries))
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	for i, j := 0, 0; i < len(queries); i++ {
		// Join all edges with distance less than d
		for j < len(edgeList) && edgeList[j][2] < queries[i][2] {
			join(edgeList[j][0], edgeList[j][1], parent, rank)
			j++
		}

		// Now check whether vertex in query are connected
		// If distance of all connected edges were lesser, it would have connected in earlier loop
		if isConnected(queries[i][0], queries[i][1], parent) {
			res[queries[i][3]] = true
		} else {
			res[queries[i][3]] = false
		}
	}

	return res
}

func join(u, v int, parent, rank []int) {
	p1 := find(u, parent)
	p2 := find(v, parent)

	if p1 == p2 {
		return
	}

	if rank[p1] > rank[p2] {
		parent[p2] = p1
	} else if rank[p2] > rank[p1] {
		parent[p1] = p2
	} else {
		parent[p2] = p1
		rank[p1]++
	}
}

func find(u int, parent []int) int {
	if u != parent[u] {
		parent[u] = find(parent[u], parent)
	}
	return parent[u]
}

func isConnected(u, v int, parent []int) bool {
	return find(parent[u], parent) == find(parent[v], parent)
}
