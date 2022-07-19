package main

import (
	"fmt"
	"leetcode_fjp/data_structures"
)

type TreeNode = data_structures.TreeNode
type ListNode = data_structures.ListNode

func matrixToString(matrix [][]int) (s string) {
	s = "\n"
	for _, m := range matrix {
		s += fmt.Sprintf("%v\n", m)
	}
	return
}

func main() {
	// mainTwoSum()
	// mainAddTwoNumbers()
	// mainReverseList()
	// mainGetLeastNumbers()
	// mainSubsetsWithDup()
	// mainMinPathSum()
	// mainRotate()
	mainTranspose()
}

/*
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
func mainReverseList() {
	nums1 := []int{1, 2, 3, 4, 5}
	l1 := data_structures.Ints2List(nums1)
	data_structures.PrintList(l1)
	l2 := reverseList(l1)
	data_structures.PrintList(l2)
}

func mainGetLeastNumbers() {
	// nums := []int{0, 1, 2, 1}
	// ret := getLeastNumbers(nums, 1)
	nums := []int{0, 0, 0, 2, 0, 5}
	ret := getLeastNumbers(nums, 0)
	fmt.Println(ret)
}
func mainSubsetsWithDup() {
	nums := []int{1, 2, 2}
	ret := subsetsWithDup(nums)
	fmt.Println(ret)
}

func mainMinPathSum() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
*/

func mainRotate() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
func rotate(matrix [][]int) {
	reverse := func(nums []int) {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		reverse(matrix[i])
	}
}

func mainTranspose() {
	// mat := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// }
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(transpose(mat))
}
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, m)
	}
	fmt.Println(m, n, ret)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// ret[i][j], ret[j][i] = matrix[j][i], matrix[i][j]
			// if j > i {
			ret[j][i] = matrix[i][j]
			// } else {
			// ret[j][i] = matrix[j][i]
			// }

		}
		fmt.Println("fjp", ret[i])
	}
	return ret
}
