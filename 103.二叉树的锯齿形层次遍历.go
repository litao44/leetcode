/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0, 1)
	if root == nil {
		return ret
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		thisLevel := make([]int, 0, 1)

		p := []*TreeNode{}

		if len(ret)%2 == 0 {

			for i := 0; i < l; i++ {
				node := queue[i]
				thisLevel = append(thisLevel, node.Val)
				if node.Left != nil {
					p = append(p, node.Left)
				}
				if node.Right != nil {
					p = append(p, node.Right)
				}
			}

		} else {

			for i := l - 1; i >= 0; i-- {
				node := queue[i]
				thisLevel = append(thisLevel, node.Val)

				if node.Right != nil {
					p = append([]*TreeNode{node.Right}, p...)
				}
				if node.Left != nil {
					p = append([]*TreeNode{node.Left}, p...)
				}
			}
		}

		queue = p
		ret = append(ret, thisLevel)
	}

	return ret
}

// @lc code=end

