package main

// 383. Ransom Note
// https://leetcode.com/problems/ransom-note/?envType=study-plan&id=data-structure-i
func canConstruct(ransomNote string, magazine string) bool {
	counter := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		counter[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		counter[ransomNote[i]-'a']--
		if counter[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
