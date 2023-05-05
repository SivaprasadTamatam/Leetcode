/*
Maximum Number of Vowels in a Substring of Given Length
iven a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

Example 1:

Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.
Example 2:

Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.
Example 3:

Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.

Constraints:

1 <= s.length <= 105
s consists of lowercase English letters.
1 <= k <= s.length
*/

package main

func maxVowels(s string, k int) int {
	n := len(s)
	maxVowels, currVowels := 0, 0
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			currVowels++
		}
	}

	maxVowels = currVowels

	for i := k; i < n; i++ {
		if vowels[s[i]] {
			currVowels++
		}
		if vowels[s[i-k]] {
			currVowels--
		}

		if currVowels > maxVowels {
			maxVowels = currVowels
		}
	}
	return maxVowels
}
