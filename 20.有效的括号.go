/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack.Push(s[i])
		} else {
			if isPair(stack.Top(), s[i]) {
				stack.Pop()
			} else {
				stack.Push(s[i])
			}
		}
	}
	return stack.IsEmpty()

}

func isPair(a, b byte) bool {
	if a == '(' && b == ')' {
		return true
	}

	if a == '{' && b == '}' {
		return true
	}

	if a == '[' && b == ']' {
		return true
	}

	return false
}

func NewStack() *Stack {
	return &Stack{
		s: make([]byte, 0),
	}
}

type Stack struct {
	s []byte
}

func (s *Stack) Push(e byte) {
	s.s = append(s.s, e)
}

func (s *Stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *Stack) Top() byte {
	if s.IsEmpty() {
		return 0
	}
	return s.s[len(s.s)-1]
}

func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	l := len(s.s)
	ret := s.s[l-1]
	s.s = s.s[:l-1]
	return ret, true
}

// @lc code=end

