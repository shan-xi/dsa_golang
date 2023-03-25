package binary_tree_preorder_traversal

import c "leetcode/common"

// 144. Binary Tree Preorder Traversal
func preorderTraversal(root *c.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*c.TreeNode{root}
	result := []int{}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return result
}
