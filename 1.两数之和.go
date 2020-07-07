/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		c := target - v
		if pos, ok := m[c]; ok {
			return []int{pos, i}
		}
		m[v] = i
	}

	return nil
}

// @lc code=end

