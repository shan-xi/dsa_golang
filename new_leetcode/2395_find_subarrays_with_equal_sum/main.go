package main

import "fmt"

func main() {
	r := findSubarrays([]int{4, 2, 4})
	fmt.Println(r)
	r = findSubarrays([]int{1, 2, 3, 4, 5})
	fmt.Println(r)
	r = findSubarrays([]int{0, 0, 0})
	fmt.Println(r)
}

func findSubarrays(nums []int) bool {
	seen := make(map[int]bool)
	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i] + nums[i+1]
		if seen[sum] {
			return true
		}
		seen[sum] = true
	}
	return false
}
