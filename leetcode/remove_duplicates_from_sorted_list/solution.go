package remove_duplicates_from_sorted_list

import c "leetcode/common"

// 83. Remove Duplicates from Sorted List
func deleteDuplicates(head *c.ListNode) *c.ListNode {
	if head == nil {
		return nil
	}
	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
