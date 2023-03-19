package reverse_linked_list

import c "leetcode/common"

// https://leetcode.com/problems/reverse-linked-list/
// 206. Reverse Linked List
func reverseList(head *c.ListNode) *c.ListNode {
	if head == nil {
		return nil
	}
	var prev *c.ListNode
	curr := head
	for curr != nil {
		tempNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = tempNode
	}
	return prev
}
