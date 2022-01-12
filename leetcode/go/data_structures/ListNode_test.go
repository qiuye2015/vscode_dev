package data_structures

import (
	"fmt"
	"testing"
)

func TestInts2List(t *testing.T) {
	// head := Ints2List([]int{})
	// head := Ints2List([]int{1})
	head := Ints2List([]int{1, 2, 3, 4, 5})
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}

func TestList2Ints(t *testing.T) {
	head := Ints2List([]int{1, 2, 3, 4, 5})
	res := List2Ints(head)
	PrintList(head)
	fmt.Println(res)
}

func TestListNode_GetNodeWith(t *testing.T) {
	head := Ints2List([]int{1, 2, 3, 4, 5})
	cur := head.GetNodeWith(2)
	if cur != nil {
		fmt.Println(cur.Val)
	}
}

func TestInts2ListWithCycle(t *testing.T) {
	ints := []int{1, 2, 3}
	l := Ints2ListWithCycle(ints, -1)
	PrintList(l)
	// PrintListWithAddr(l, 4)
	fmt.Println("Cycle")
	l = Ints2ListWithCycle(ints, 0)
	PrintListWithAddr(l, 4)
}
