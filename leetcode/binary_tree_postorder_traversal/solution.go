package binary_tree_postorder_traversal

import c "leetcode/common"

var result []int = []int{}

// 145. Binary Tree Postorder Traversal
func postorderTraversal(root *c.TreeNode) []int {
	result = []int{}
	postOrder(root)
	return result
}

func postOrder(root *c.TreeNode) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	result = append(result, root.Val)
}
