/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	return (l.Val == r.Val) && helper(l.Right, r.Left) && helper(l.Left, r.Right)
}

func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root, root}
	for len(queue) > 0 {
		l, r := queue[0], queue[1]
		queue = queue[2:]
		if l == nil && r == nil {
			continue
		}

		if l == nil || r == nil {
			return false
		}

		if l.Val != r.Val {
			return false
		}

		queue = append(queue, l.Left, r.Right, l.Right, r.Left)
	}

	return true
}

// @lc code=end

