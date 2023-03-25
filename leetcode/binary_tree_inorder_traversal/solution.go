package binary_tree_inorder_traversal

import c "leetcode/common"

var result []int = []int{}

// 94. Binary Tree Inorder Traversal
func inorderTraversal(root *c.TreeNode) []int {
	result = []int{}
	inOrder(root)
	return result
}

func inOrder(root *c.TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	result = append(result, root.Val)
	inOrder(root.Right)
}
