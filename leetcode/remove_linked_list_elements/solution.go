package remove_linked_list_elements

import c "leetcode/common"

// 203. Remove Linked List Elements
func removeElements(head *c.ListNode, val int) *c.ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
