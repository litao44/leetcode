/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	stack stack
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {

}

func (ms *MinStack) Push(x int) {
	ms.stack.Push(x)
	if x < ms.min {
		ms.min = x
	}
}

func (ms *MinStack) Pop() {

}

func (ms *MinStack) Top() int {

}

func (ms *MinStack) GetMin() int {

}

type stack struct {
	data []int
}

func (s *stack) Push(e int) {
	s.data = append(s.data, e)
}

func (s *stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) Top() int {
	if s.IsEmpty() {
		return 0
	}
	return s.data[len(s.s)-1]
}

func (s *stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	l := len(s.data)
	ret := s.data[l-1]
	s.s = s.data[:l-1]
	return ret, true
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

