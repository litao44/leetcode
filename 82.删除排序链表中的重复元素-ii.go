/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
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

	distinct := make(map[int]int, 1)
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := distinct[cur.Val]; ok {
			distinct[cur.Val]++
		} else {
			distinct[cur.Val] = 1
		}
	}

	ret := &ListNode{}
	for cur := head; cur != nil; cur = cur.Next {
		if distinct[cur.Val] == 1 {
			ret.Next = cur
			ret = ret.Next
		}
	}

	return ret.Next
}

// @lc code=end

