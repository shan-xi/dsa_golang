package leetcode

// 101. Symmetric Tree
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(leftNode *TreeNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}
	if leftNode == nil || rightNode == nil {
		return false
	}
	if leftNode.Val != rightNode.Val {
		return false
	}
	return isSymmetricHelper(leftNode.Left, rightNode.Right) && isSymmetricHelper(leftNode.Right, rightNode.Left)
}
