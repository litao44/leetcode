/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	p, p1, p2 := m+n-1, m-1, n-1
	for ; p1 >= 0 && p2 >= 0; p-- {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
	}

	copy(nums1, nums2[:p2+1])
}

// @lc code=end

