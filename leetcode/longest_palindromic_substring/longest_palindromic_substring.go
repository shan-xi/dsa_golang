package main

// func longestPalindrome(s string) string {
// 	if len(s) == 1 {
// 		return s
// 	}
// 	sa := []rune(s)
// 	trace := [128]int{}
// 	for i := 0; i < len(trace); i++ {
// 		trace[i] = -1
// 	}

// 	right := 0
// 	left := 0
// 	max_right := 0
// 	max_left := 0
// 	for right < len(sa) {
// 		if trace[sa[right]] == -1 {
// 			trace[sa[right]] = right
// 		} else {
// 			left = trace[sa[right]]
// 			for right-left > max_right-max_left {
// 				if isP(sa[left : right+1]) {
// 					// fmt.Printf("%c\n", sa[left:right+1])
// 					if max_right-max_left < right-left {
// 						max_right = right
// 						max_left = left
// 					}
// 				}
// 				left += 1
// 			}
// 		}
// 		right += 1
// 	}
// 	return string(sa[max_left : max_right+1])
// }

// func isP(s []rune) bool {
// 	// fmt.Printf("%c\n", s)
// 	i := 0
// 	j := len(s) - 1
// 	for i < j {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 		i += 1
// 		j -= 1
// 	}
// 	return true
// }

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	sa := []rune(s)
	start := 0
	end := 0
	for i := 0; i < len(sa); i++ {
		len1 := expandAroundCenter(sa, i, i)
		len2 := expandAroundCenter(sa, i, i+1)
		len := max(len1, len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return string(sa[start : end+1])
}

func expandAroundCenter(sa []rune, left int, right int) int {
	L := left
	R := right
	for L >= 0 && R < len(sa) && sa[L] == sa[R] {
		L -= 1
		R += 1
	}
	return R - L - 1
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
