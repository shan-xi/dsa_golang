package main

// 509. Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/?envType=study-plan&id=dynamic-programming-i
func fib(n int) int {
	first := 0
	second := 1
	if n == 0 {
		return first
	}
	if n == 1 {
		return second
	}

	for i := 2; i <= n; i++ {
		temp := second
		second = first + second
		first = temp
	}
	return second
}
