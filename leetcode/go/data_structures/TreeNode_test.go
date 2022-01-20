package data_structures

import (
	"fmt"
	"testing"
)

var (
	// 同一个 TreeNode 的不同表达方式
	//            1
	//      	/  \
	//         2    3
	//        / \  /  \
	//       4  5  6   7
	LeetCodeOrder = []int{1, 2, 3, 4, 5, 6, 7}
	preOrder      = []int{1, 2, 4, 5, 3, 6, 7}
	inOrder       = []int{4, 2, 5, 1, 6, 3, 7}
	postOrder     = []int{4, 5, 2, 6, 7, 3, 1}
)

func TestPrintTree(t *testing.T) {
	// root := Ints2Tree(LeetCodeOrder)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	// nums := []int{1, 2, 3, -1, 5, 6, 7}
	root := Ints2Tree(nums)
	// PrintTree(root)
	PrintBinaryTreeGraph(root)
}

func TestInts2Tree(t *testing.T) {
	root := Ints2Tree(LeetCodeOrder)

	pre := Tree2PreOrder(root)
	fmt.Println(pre)
	in := Tree2InOrder(root)
	fmt.Println(in)
	post := Tree2PostOrder(root)
	fmt.Println(post)

}

func TestTree2Ints(t *testing.T) {
	root := Ints2Tree(LeetCodeOrder)
	res := Tree2Ints(root)
	fmt.Println(res)
}

// func TestTree2PreOrder(t *testing.T) {
// }

// func TestTree2InOrder(t *testing.T) {
// }

// func TestTree2PostOrder(t *testing.T) {
// }
