package main

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	sa := []rune(s)
	trace := [128]int{}
	for i := 0; i < len(trace); i++ {
		trace[i] = -1
	}

	right := 0
	left := 0
	max_right := 0
	max_left := 0
	for right < len(sa) {
		if trace[sa[right]] == -1 {
			trace[sa[right]] = right
		} else {
			left = trace[sa[right]]
			for right-left > max_right-max_left {
				if isP(sa[left : right+1]) {
					// fmt.Printf("%c\n", sa[left:right+1])
					if max_right-max_left < right-left {
						max_right = right
						max_left = left
					}
				}
				left += 1
			}
		}
		right += 1
	}
	return string(sa[max_left : max_right+1])
}

func isP(s []rune) bool {
	// fmt.Printf("%c\n", s)
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i += 1
		j -= 1
	}
	return true
}
