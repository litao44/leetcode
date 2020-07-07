/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	size := len(nums)
	switch size {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	first, second := nums[0], max(nums[1], nums[0])
	for index := 2; index < size; index++ {
		temp := second
		second = max(nums[index]+first, second)
		first = temp
	}

	return second
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

// @lc code=end

