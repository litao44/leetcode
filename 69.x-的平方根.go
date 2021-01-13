/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	lo, hi := 0, x
	ans := -1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid*mid <= x {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return ans
}

// @lc code=end

