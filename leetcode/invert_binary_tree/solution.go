package leetcode

// 226. Invert Binary Tree
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTreeHelper(root)
	return root
}

func invertTreeHelper(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	invertTreeHelper(root.Left)
	invertTreeHelper(root.Right)
}
