package merge_two_sorted_lists

import c "leetcode/common"

// 21. Merge Two Sorted Lists
func mergeTwoLists(list1 *c.ListNode, list2 *c.ListNode) *c.ListNode {
	dummy := c.ListNode{0, nil}
	currNode := &dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			currNode.Next = list1
			list1 = list1.Next
		} else {
			currNode.Next = list2
			list2 = list2.Next
		}
		currNode = currNode.Next
	}
	if list1 != nil {
		currNode.Next = list1
	}
	if list2 != nil {
		currNode.Next = list2
	}
	return dummy.Next
}
