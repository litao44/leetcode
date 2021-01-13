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
	return recursion(root)
}

func recursion(root *TreeNode) []int {
	m := map[int]int{}
	recursionHelper(root, m, 0)

	result := make([]int, len(m))
	for key, val := range m {
		result[key] = val
	}

	return result
}

func recursionHelper(root *TreeNode, result map[int]int, curentDepth int) {
	if root == nil {
		return
	}

	if _, ok := result[curentDepth]; !ok {
		result[curentDepth] = root.Val
	} else {
		result[curentDepth] = max(result[curentDepth], root.Val)
	}

	recursionHelper(root.Right, result, curentDepth+1)
	recursionHelper(root.Left, result, curentDepth+1)
}

func bfs(root *TreeNode) []int {
	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

