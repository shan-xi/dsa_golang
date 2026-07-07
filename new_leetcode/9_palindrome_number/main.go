package main

import "fmt"

func main() {
	r := isPalindrome(121)
	fmt.Println(r)

	r = isPalindrome(-121)
	fmt.Println(r)

	r = isPalindrome(10)
	fmt.Println(r)
}

func isPalindrome(x int) bool {
	original := x
	reversed := 0

	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	return original == reversed
}
