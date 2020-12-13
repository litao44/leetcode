/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	root := head
	for root.Next != nil {
		if root.Val == root.Next.Val {
			root.Next = root.Next.Next
		} else {
			root = root.Next
		}

	}
	return head

}

// @lc code=end

