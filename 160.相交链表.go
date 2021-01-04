/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	store := make(map[*ListNode]bool, 1)

	for iter := headA; iter != nil; iter = iter.Next {
		store[iter] = true
	}

	for iter := headB; iter != nil; iter = iter.Next {
		if _, ok := store[iter]; ok {
			return iter
		}
	}

	return nil
}

// @lc code=end

