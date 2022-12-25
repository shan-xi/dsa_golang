package main

// import "fmt"

// func lengthOfLongestSubstring(s string) int {
// 	sArr := []rune(s)
// 	checkExist := make(map[rune]bool)
// 	indexMap := make(map[rune]int)
// 	pos := 0
// 	temp := 0
// 	max := 0
// 	for pos < len(sArr) {
// 		c := sArr[pos]
// 		if !checkExist[c] {
// 			checkExist[c] = true
// 			indexMap[c] = pos
// 			temp++
// 			// fmt.Printf("%c %d\n", c, temp)
// 		} else {
// 			// fmt.Printf("repeat! %c %d %d\n", c, pos, indexMap[c])
// 			if temp > max {
// 				max = temp
// 			}
// 			// checkExist = make(map[rune]bool)
// 			clear(checkExist)
// 			pos = indexMap[c]
// 			temp = 0
// 		}
// 		pos += 1
// 	}
// 	if temp > max {
// 		max = temp
// 	}

// 	return max
// }

// func clear(m map[rune]bool) {
// 	for k := range m {
// 		m[k] = false
// 	}
// }

// func lengthOfLongestSubstring(s string) int {
// 	sa := []rune(s)
// 	chars := make(map[rune]int)
// 	left := 0
// 	right := 0
// 	res := 0
// 	for right < len(sa) {
// 		r := sa[right]
// 		chars[r] += 1
// 		for chars[r] > 1 {
// 			l := sa[left]
// 			chars[l] -= 1
// 			left += 1
// 		}

// 		res = max(res, (right - left + 1))
// 		right += 1
// 	}
// 	return res
// }
// func lengthOfLongestSubstring(s string) int {
// 	sa := []rune(s)
// 	n := len(s)
// 	ans := 0
// 	mp := make(map[rune]int)
// 	i := 0
// 	for j := 0; j < n; j++ {
// 		if v, ok := mp[sa[j]]; ok {
// 			i = max(v, i)
// 		}

// 		ans = max(ans, j-i)
// 		mp[sa[j]] = j
// 	}
// 	return ans
// }

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring(s string) int {
	sa := []rune(s)
	chars := [128]int{}
	for i := 0; i < len(chars); i++ {
		chars[i] = -1
	}
	left := 0
	right := 0
	res := 0
	for right < len(s) {
		r := sa[right]
		index := chars[int(r)]
		if index != -1 && left <= index && index < right {
			left = index + 1
		}
		res = max(res, right-left+1)
		chars[int(r)] = right
		right += 1
	}
	return res
}
