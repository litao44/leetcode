/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}

	return maxEle(dp)
}

func maxEle(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

