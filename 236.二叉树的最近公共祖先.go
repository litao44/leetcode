/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//1. 记录父节点
	parent := make(map[int]*TreeNode)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			parent[node.Left.Val] = node
			dfs(node.Left)
		}

		if node.Right != nil {
			parent[node.Right.Val] = node
			dfs(node.Right)
		}
	}

	dfs(root)

	visitied := make(map[int]struct{})
	for p != nil {
		visitied[p.Val] = struct{}{}
		p = parent[p.Val]
	}

	for q != nil {
		if _, ok := visitied[q.Val]; ok {
			return q
		}
		q = parent[q.Val]
	}

	return nil
}

// @lc code=end

