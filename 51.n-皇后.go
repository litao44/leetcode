/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = repeatDot(n)
	}

	result := make([][]string, 0, 1)
	backtrack(&result, board, 0)
	return result
}

func backtrack(result *[][]string, board []string, row int) {
	if row == len(board) {
		temp := make([]string, len(board))
		copy(temp, board)
		*result = append(*result, temp)
		return
	}

	n := len(board[row])
	for col := 0; col < n; col++ {
		if !isValid(board, row, col) {
			continue
		}

		b := []byte(board[row])
		b[col] = 'Q'
		board[row] = string(b)
		backtrack(result, board, row+1)
		b[col] = '.'
		board[row] = string(b)
	}
}

func isValid(board []string, row, col int) bool {
	n := len(board)

	// 列
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 右上角
	i, j := row-1, col+1
	for i >= 0 && j < n {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}

	// 左上角
	i, j = row-1, col-1
	for i >= 0 && j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	return true
}

func repeatDot(n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += "."
	}
	return res
}

// @lc code=end

