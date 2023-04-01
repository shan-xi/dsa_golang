package leetcode

import "testing"

func Test_case1_success(t *testing.T) {
	n1 := &TreeNode{4, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{7, nil, nil}
	n4 := &TreeNode{1, nil, nil}
	n5 := &TreeNode{3, nil, nil}
	n6 := &TreeNode{6, nil, nil}
	n7 := &TreeNode{9, nil, nil}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	a := invertTree(n1)
	if n3 != a.Left {
		t.Errorf("n3 %v != a.Left %v", n3, a.Left)
	}
	if n7 != a.Left.Left {
		t.Errorf("n7 %v != a.Left.Left %v", n7, a.Left.Left)
	}
	if n6 != a.Left.Right {
		t.Errorf("n6 %v != a.Left.Left %v", n6, a.Left.Left)
	}
	if n2 != a.Right {
		t.Errorf("n2 %v != a.Right %v", n2, a.Right)
	}
	if n5 != a.Right.Left {
		t.Errorf("n5 %v != a.Right.Left %v", n5, a.Right.Left)
	}
	if n4 != a.Right.Right {
		t.Errorf("n4 %v != a.Right.Right %v", n4, a.Right.Right)
	}
}

func Test_case2_success(t *testing.T) {
	n1 := &TreeNode{2, nil, nil}
	n2 := &TreeNode{1, nil, nil}
	n3 := &TreeNode{3, nil, nil}
	n1.Left = n2
	n1.Right = n3
	a := invertTree(n1)
	if n3 != a.Left {
		t.Errorf("n3 %v != a.Left %v", n3, a.Left)
	}
	if n2 != a.Right {
		t.Errorf("n2 %v != a.Right %v", n2, a.Right)
	}
}

func Test_case3_success(t *testing.T) {
	a := invertTree(nil)
	if a != nil {
		t.Errorf("nil %v != a %v", nil, a)
	}
}
