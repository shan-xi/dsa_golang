package merge_two_sorted_lists

import (
	c "leetcode/common"
	"testing"
)

func Test_case1_success(t *testing.T) {
	n1 := c.ListNode{0, nil}
	var list1 c.ListNode
	var list2 c.ListNode
	list2 = n1

	a := mergeTwoLists(&list1, &list2)
	e := &n1
	for e != nil && a != nil {
		e = e.Next
		a = a.Next
		if e != a {
			t.Errorf("e %v != a %v", e, a)
		}
	}
}
