package main

// 387. First Unique Character in a String
// https://leetcode.com/problems/first-unique-character-in-a-string/description/?envType=study-plan&id=data-structure-i
func firstUniqChar(s string) int {
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a'] += 1
	}
	for i := 0; i < len(s); i++ {
		if count[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
