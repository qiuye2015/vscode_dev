package data_structures

type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}
func (q *Queue) Pop() int {
	ret := q.nums[0]
	q.nums = q.nums[1:]
	return ret
}
func (q *Queue) Len() int {
	return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
