package binary_tree_inorder_traversal

import c "leetcode/common"

// 94. Binary Tree Inorder Traversal
func inorderTraversal(root *c.TreeNode) []int {
	var result []int
	var stack []*c.TreeNode
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)
		current = current.Right
	}
	return result
}
