package data_structures

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	if head == nil {
		return
	}
	limit := 100
	times := 0

	cur := head
	for cur != nil {
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}
		if cur.Next != nil {
			fmt.Printf("%v -> ", cur.Val)
		} else {
			fmt.Printf("%v -> nil\n", cur.Val)
		}
		cur = cur.Next
	}
}

func PrintListWithAddr(head *ListNode, limit int) {
	if head == nil {
		return
	}
	times := 0
	cur := head
	for cur != nil {
		times++
		if times > limit {
			fmt.Printf("链条深度超过%d，可能出现环状链条。请检查错误。", limit)
			return
		}
		// if cur.Next != nil {

		fmt.Printf("[%v %p %p] -> \n", cur.Val, cur, cur.Next)
		// } else {
		// fmt.Printf("%v -> nil\n", cur.Val)
		// }
		cur = cur.Next
	}
}

func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	fakerHead := &ListNode{}
	cur := fakerHead
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return fakerHead.Next
}

func List2Ints(head *ListNode) []int {
	res := []int{}
	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

// GetNodeWith returns the first node with val
func (l *ListNode) GetNodeWith(val int) *ListNode {
	cur := l
	for cur != nil {
		if cur.Val == val {
			break
		}
		cur = cur.Next
	}
	return cur
}

// Ints2ListWithCycle returns a list whose tail point to pos-indexed node
// head's index is 0
// if pos = -1, no cycle
// 数组转为带环链表
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}
