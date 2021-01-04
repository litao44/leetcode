/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	matche := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if matche(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if matche(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

// @lc code=end

