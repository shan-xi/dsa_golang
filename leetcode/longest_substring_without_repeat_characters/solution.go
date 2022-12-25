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

func lengthOfLongestSubstring(s string) int {
	sa := []rune(s)
	chars := make(map[rune]int)
	left := 0
	right := 0
	res := 0
	for right < len(sa) {
		r := sa[right]
		chars[r] += 1
		for chars[r] > 1 {
			l := sa[left]
			chars[l] -= 1
			left += 1
		}

		res = max(res, (right - left + 1))
		right += 1
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
