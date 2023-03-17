package main

// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/?envType=study-plan&id=data-structure-i
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		count[t[i]-'a']--
	}

	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true
}
