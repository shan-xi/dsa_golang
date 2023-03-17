package main

// 383. Ransom Note
// https://leetcode.com/problems/ransom-note/?envType=study-plan&id=data-structure-i
func canConstruct(ransomNote string, magazine string) bool {
	countChars := make([]int, 128)
	for _, c := range magazine {
		countChars[c]++
	}
	for _, c := range ransomNote {
		if countChars[c] == 0 {
			return false
		}
		countChars[c]--
	}
	return true
}
