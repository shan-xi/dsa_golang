package leetcode

// 104. Maximum Depth of Binary Tree
func maxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, d int) int {
	if root == nil {
		return d
	}
	left := dfs(root.Left, d+1)
	right := dfs(root.Right, d+1)
	if left > right {
		return left
	} else {
		return right
	}
}
