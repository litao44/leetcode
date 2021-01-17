/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	next := getNext(needle)
	i, j := 0, 0

	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == len(needle) {
		return i - j
	}
	return -1
}

func getNext(s string) []int {
	next := make([]int, len(s)+1)
	next[0] = -1

	i, j := 0, -1
	for i < len(s) {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

// @lc code=end

