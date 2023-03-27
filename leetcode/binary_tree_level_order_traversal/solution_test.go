package leetcode

import "testing"

func Test_case1_success(t *testing.T) {
	n1 := &TreeNode{3, nil, nil}
	n2 := &TreeNode{9, nil, nil}
	n3 := &TreeNode{20, nil, nil}
	n4 := &TreeNode{15, nil, nil}
	n5 := &TreeNode{7, nil, nil}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	a := levelOrder(n1)
	e := [][]int{{3}, {9, 20}, {15, 7}}
	for i := range a {
		for j := range a[i] {
			if e[i][j] != a[i][j] {
				t.Errorf("e %v != a %v", e, a)
			}
		}
	}
}
