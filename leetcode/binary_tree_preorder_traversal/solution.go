package binary_tree_preorder_traversal

import c "leetcode/common"

var result []int = []int{}

// 144. Binary Tree Preorder Traversal
func preorderTraversal(root *c.TreeNode) []int {
	result = []int{}
	preOrder(root)
	return result
}

func preOrder(root *c.TreeNode) {
	if root == nil {
		return
	}
	result = append(result, root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}
