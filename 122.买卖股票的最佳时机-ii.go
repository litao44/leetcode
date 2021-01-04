/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	maxPro := 0
	for i := 1; i < len(prices); i++ {
		if profit := prices[i] - prices[i-1]; profit > 0 {
			maxPro += profit
		}
	}

	return maxPro
}

// @lc code=end

