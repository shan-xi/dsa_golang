package main

// 141. Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	first := head
	second := head
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
		if first == second {
			return true
		}
	}
	return false
}
