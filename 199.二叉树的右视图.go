/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := []int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		l := len(queue)
		thisLevel := make([]int, 0, l)
		for i := 0; i < l; i++ {
			node := queue[i]
			thisLevel = append(thisLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, thisLevel[len(thisLevel)-1])
		queue = queue[l:]
	}

	return ans
}

// @lc code=end

