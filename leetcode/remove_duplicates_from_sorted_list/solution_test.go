package remove_duplicates_from_sorted_list

import (
	c "leetcode/common"
	"strconv"
	"testing"
)

func Test_case1_success(t *testing.T) {
	n1 := &c.ListNode{1, nil}
	n2 := &c.ListNode{1, nil}
	n3 := &c.ListNode{2, nil}
	n1.Next = n2
	n2.Next = n3
	a := deleteDuplicates(n1)
	str := ""
	for a != nil {
		if str != "" {
			str += ", " + strconv.Itoa(a.Val)
		} else {
			str += strconv.Itoa(a.Val)
		}
		a = a.Next
	}
	str = "[" + str + "]"
	e := "[1, 2]"
	if e != str {
		t.Errorf("actual %v != e %v ", str, e)
	}
}
func Test_case2_success(t *testing.T) {
	n1 := &c.ListNode{1, nil}
	n2 := &c.ListNode{1, nil}
	n3 := &c.ListNode{2, nil}
	n4 := &c.ListNode{3, nil}
	n5 := &c.ListNode{3, nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	a := deleteDuplicates(n1)
	str := ""
	for a != nil {
		if str != "" {
			str += ", " + strconv.Itoa(a.Val)
		} else {
			str += strconv.Itoa(a.Val)
		}
		a = a.Next
	}
	str = "[" + str + "]"
	e := "[1, 2, 3]"
	if e != str {
		t.Errorf("actual %v != e %v ", str, e)
	}
}
func Test_case3_success(t *testing.T) {
	n1 := &c.ListNode{1, nil}
	a := deleteDuplicates(n1)
	str := ""
	for a != nil {
		if str != "" {
			str += ", " + strconv.Itoa(a.Val)
		} else {
			str += strconv.Itoa(a.Val)
		}
		a = a.Next
	}
	str = "[" + str + "]"
	e := "[1]"
	if e != str {
		t.Errorf("actual %v != e %v ", str, e)
	}
}
