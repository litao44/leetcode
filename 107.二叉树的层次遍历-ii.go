/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {

	ret := make([][]int, 0, 1)
	if root == nil {
		return ret
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		thisLevel := make([]int, 0, 1)
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

		queue = queue[l:]
		ret = append([][]int{thisLevel}, ret...)
	}

	return ret
}

// @lc code=end

