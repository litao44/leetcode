/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 || l2 == 0 {
		return l1 + l2
	}

	dp := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}

	for j := 0; j < l2+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
			}
		}
	}

	return dp[l1][l2]
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}
	return minNum
}

// @lc code=end

