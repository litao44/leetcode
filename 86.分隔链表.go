/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallP := small
	large := &ListNode{}
	largeP := large

	for ; head != nil; head = head.Next {
		if head.Val < x {
			smallP.Next = head
			smallP = smallP.Next
		} else {
			largeP.Next = head
			largeP = largeP.Next
		}
	}

	largeP.Next = nil
	smallP.Next = large.Next
	return small.Next
}

// @lc code=end

