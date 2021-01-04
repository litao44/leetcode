/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var rHead, rTail *ListNode = head, head
	for i := 0; i < k; i++ {
		if rTail == nil {
			return head
		}
		rTail = rTail.Next
	}

	newHead := reverse(rHead, rTail)
	rHead.Next = reverseKGroup(rTail, k)
	return newHead
}

func reverse(head, tail *ListNode) *ListNode {
	var cur, pre *ListNode = head, nil
	for cur != tail {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

// @lc code=end

