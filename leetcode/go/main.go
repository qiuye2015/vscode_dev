package main

import (
	"fmt"
	"leetcode_fjp/data_structures"
)

type TreeNode = data_structures.TreeNode
type ListNode = data_structures.ListNode

func main() {
	// mainTwoSum()
	// mainAddTwoNumbers()
}

//---------
func mainTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}

func mainAddTwoNumbers() {
	nums1 := []int{9, 9, 9, 9, 9, 9, 9}
	nums2 := []int{9, 9, 9, 9}
	l1 := data_structures.Ints2List(nums1)
	l2 := data_structures.Ints2List(nums2)
	data_structures.PrintList(l1)
	data_structures.PrintList(l2)
	l3 := addTwoNumbers(l1, l2)
	data_structures.PrintList(l3)
}
