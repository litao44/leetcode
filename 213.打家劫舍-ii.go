/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	size := len(nums)
	switch size {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}

	return max(rob2(nums[1:]), rob2(nums[:len(nums)-1]))
}

func rob2(nums []int) int {
	first, second := nums[0], max(nums[1], nums[0])
	for index := 2; index < len(nums); index++ {
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

