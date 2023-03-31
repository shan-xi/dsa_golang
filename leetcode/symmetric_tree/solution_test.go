package leetcode

import "testing"

func Test_case1_success(t *testing.T) {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{2, nil, nil}
	n4 := &TreeNode{3, nil, nil}
	n5 := &TreeNode{4, nil, nil}
	n6 := &TreeNode{4, nil, nil}
	n7 := &TreeNode{3, nil, nil}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	a := isSymmetric(n1)
	e := true
	if e != a {
		t.Errorf("e %v != a %v", e, a)
	}
}

func Test_case2_success(t *testing.T) {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{2, nil, nil}
	n4 := &TreeNode{3, nil, nil}
	n5 := &TreeNode{3, nil, nil}
	n1.Left = n2
	n1.Right = n3
	n2.Right = n4
	n3.Right = n5
	a := isSymmetric(n1)
	e := false
	if e != a {
		t.Errorf("e %v != a %v", e, a)
	}
}
