/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start

func NewStack() *Stack {
	return &Stack{
		s: make([]int, 0),
	}
}

type Stack struct {
	s []int
}

func (s *Stack) Push(e int) {
	s.s = append(s.s, e)
}

func (s *Stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		return 0
	}
	return s.s[len(s.s)-1]
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	l := len(s.s)
	ret := s.s[l-1]
	s.s = s.s[:l-1]
	return ret, true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func longestValidParentheses(s string) int {
	var maxLen int
	stack := NewStack()
	stack.Push(-1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				maxLen = max(maxLen, i-stack.Top())
			}
		}
	}
	return maxLen
}

// @lc code=end

