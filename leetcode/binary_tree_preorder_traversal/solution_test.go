package binary_tree_preorder_traversal

import (
	c "leetcode/common"
	"testing"
)

func Test_case1_success(t *testing.T) {
	n1 := &c.TreeNode{1, nil, nil}
	n2 := &c.TreeNode{2, nil, nil}
	n3 := &c.TreeNode{3, nil, nil}
	n1.Right = n2
	n2.Left = n3
	a := preorderTraversal(n1)
	e := []int{1, 2, 3}
	for i := range a {
		if e[i] != a[i] {
			t.Errorf("e %v != a %v", e[i], a[i])
		}
	}
}

func Test_case2_success(t *testing.T) {
	n1 := &c.TreeNode{1, nil, nil}
	a := preorderTraversal(n1)
	e := []int{1}
	for i := range a {
		if e[i] != a[i] {
			t.Errorf("e %v != a %v", e[i], a[i])
		}
	}
}

func Test_case3_success(t *testing.T) {
	a := preorderTraversal(nil)
	e := []int{}
	for i := range a {
		if e[i] != a[i] {
			t.Errorf("e %v != a %v", e[i], a[i])
		}
	}
}
