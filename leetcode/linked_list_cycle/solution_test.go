package main

import "testing"

func Test_case1_success(t *testing.T) {
	n4 := ListNode{-4, nil}
	n3 := ListNode{0, &n4}
	n2 := ListNode{2, &n3}
	n1 := ListNode{3, &n2}
	n4.Next = &n2

	a := hasCycle(&n1)
	e := true
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}

func Test_case2_success(t *testing.T) {
	n2 := ListNode{2, nil}
	n1 := ListNode{1, &n2}

	a := hasCycle(&n1)
	e := false
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}

func Test_case3_success(t *testing.T) {
	n2 := ListNode{2, nil}
	n1 := ListNode{1, &n2}
	n2.Next = &n1

	a := hasCycle(&n1)
	e := true
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}

func Test_case4_success(t *testing.T) {
	n1 := ListNode{1, nil}

	a := hasCycle(&n1)
	e := false
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}

func Test_case5_success(t *testing.T) {
	n4 := ListNode{1, nil}
	n3 := ListNode{1, nil}
	n2 := ListNode{1, nil}
	n1 := ListNode{1, nil}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4

	a := hasCycle(&n1)
	e := false
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}
