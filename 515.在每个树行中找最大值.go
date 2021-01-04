/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
func largestValues(root *TreeNode) []int {
	res := make([]int, 0, 1)
	helper(root, 0, res)
	return res
}

func helper(root *TreeNode, depth int, res []int) {
	if root == nil {
		return
	}

	currentDepth := depth + 1

}

// @lc code=end

