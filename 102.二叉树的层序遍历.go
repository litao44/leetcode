/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0, 1)
	if root == nil {
		return ret
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		thisLevelLen := len(queue)
		thisLevel := make([]int, 0, 1)
		for i := 0; i < thisLevelLen; i++ {
			node := queue[i]
			thisLevel = append(thisLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[thisLevelLen:]
		ret = append(ret, thisLevel)
	}

	return ret
}

// @lc code=end

