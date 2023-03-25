package binary_tree_postorder_traversal

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
	a := postorderTraversal(n1)
	e := []int{3, 2, 1}
	for i := range a {
		if e[i] != a[i] {
			t.Errorf("e %v != a %v", e[i], a[i])
		}
	}
}

func Test_case2_success(t *testing.T) {
	n1 := &c.TreeNode{1, nil, nil}
	a := postorderTraversal(n1)
	e := []int{1}
	for i := range a {
		if e[i] != a[i] {
			t.Errorf("e %v != a %v", e[i], a[i])
		}
	}
}

func Test_case3_success(t *testing.T) {
	a := postorderTraversal(nil)
	e := []int{}
	for i := range a {
		if e[i] != a[i] {
			t.Errorf("e %v != a %v", e[i], a[i])
		}
	}
}
