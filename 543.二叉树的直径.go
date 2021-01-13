/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := dfs(node.Left)
		r := dfs(node.Right)

		ans = max(r+l+1, ans)
		return max(l, r) + 1
	}

	dfs(root)

	return ans - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

