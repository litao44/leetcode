/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	data *stack
	min  *stack
}

const INT_MAX = int(^uint(0) >> 1)

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: &stack{make([]int, 0, 1)},
		min:  &stack{[]int{INT_MAX}},
	}
}

func (ms *MinStack) Push(x int) {
	ms.data.Push(x)
	ms.min.Push(min(x, ms.min.Top()))
}

func (ms *MinStack) Pop() {
	ms.data.Pop()
	ms.min.Pop()
}

func (ms *MinStack) Top() int {
	return ms.data.Top()
}

func (ms *MinStack) GetMin() int {
	return ms.min.Top()
}

type stack struct {
	data []int
}

func (s *stack) Push(e int) {
	s.data = append(s.data, e)
}

func (s *stack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *stack) Pop() (int, bool) {
	l := len(s.data)
	ret := s.data[l-1]
	s.data = s.data[:l-1]
	return ret, true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

