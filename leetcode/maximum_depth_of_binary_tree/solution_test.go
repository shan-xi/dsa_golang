package leetcode

import "testing"

func Test_case1_success(t *testing.T) {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{9, nil, nil}
	n3 := &TreeNode{20, nil, nil}
	n4 := &TreeNode{15, nil, nil}
	n5 := &TreeNode{7, nil, nil}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	a := maxDepth(n1)
	e := 3
	if e != a {
		t.Errorf("e %v != a %v", e, a)
	}
}

func Test_case2_success(t *testing.T) {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n1.Right = n2
	a := maxDepth(n1)
	e := 2
	if e != a {
		t.Errorf("e %v != a %v", e, a)
	}
}
