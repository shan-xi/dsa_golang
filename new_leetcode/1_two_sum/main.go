package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	r := twoSum(nums, target)
	fmt.Println(r)

	nums = []int{3, 2, 4}
	target = 6
	r = twoSum(nums, target)
	fmt.Println(r)
}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		compliment := target - num
		if idx, ok := seen[compliment]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}
