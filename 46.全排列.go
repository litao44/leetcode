/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	result := make([][]int, 0, permuteNum(len(nums)))

	track := make([]int, 0, len(nums))
	backtrack(&result, nums, track)
	return result
}

func backtrack(result *[][]int, nums []int, track []int) {
	if len(nums) == len(track) {
		temp := make([]int, len(track))
		copy(temp, track)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if numInArr(nums[i], track) {
			continue
		}
		track = append(track, nums[i])
		backtrack(result, nums, track)
		track = track[:len(track)-1]
	}
}

func numInArr(num int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if num == arr[i] {
			return true
		}
	}

	return false
}

func permuteNum(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// @lc code=end

