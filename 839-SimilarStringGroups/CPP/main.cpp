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
#include<iostream>
#include<vector>
using namespace std;

bool CheckIsSimilar(string s1, string s2){
    if(s1 == s2) return true; // two strings X and Y are similar if they are equal.

    int diff = 0;
    for (int i = 0; i< s1.length(); i++){
        if(s1[i] != s2[i]){
            ++diff;
        }
    }
    return diff == 2; // Two strings X and Y are similar if we can swap two letters (in different positions) of X, so that it equals Y.
}
void dfs(int cp, vector<bool>& visited, vector<vector<int>>& adjList){
    visited[cp] = true;
    for(int np : adjList[cp]){
        if(!visited[np]){
            dfs(np, visited, adjList);
        }
    }
}

int numSimilarGroups(vector<string>& strs) {
    int ans = 0, n = strs.size();
    vector<vector<int>> adjList(n);

    for(int i =0; i<n; i++){
        for(int j = i+1; j<n; j++){
            if(CheckIsSimilar(strs[i], strs[j])){
                adjList[i].push_back(j);
                adjList[j].push_back(i);
            }
        }
    }

    vector<bool> visited(n, false);
    for( int i =0; i<n; i++){
        if(!visited[i]){
            ans++;
            dfs(i, visited,adjList);
        }
    }
    return ans;
}

int main(){
    vector<string> strs = {"tars","rats","arts","star"};
    cout<<numSimilarGroups(strs)<<endl;
     vector<string> strs2 = {"omv","ovm"};
    cout<<numSimilarGroups(strs2)<<endl;
}