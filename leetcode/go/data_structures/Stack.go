package data_structures

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{nums: make([]int, 0)}
}

func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

func (s *Stack) Pop() int {
	length := len(s.nums)
	ret := s.nums[length-1]
	s.nums = s.nums[:length-1]
	return ret
}

func (s *Stack) Len() int {
	return len(s.nums)
}

func (s *Stack) IsEmpty() bool {
	return len(s.nums) == 0
}
