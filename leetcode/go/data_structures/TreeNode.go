package data_structures

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL 方便添加测试数据
var NULL = -1 << 63

func PrintTree(root *TreeNode) {
	queue := []*TreeNode{root}
	n := 3
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if i == 0 {
				for j := 0; j < n; j++ {
					fmt.Printf("\t")
				}
				n--
			}
			node := queue[i]
			if node == nil {
				fmt.Printf("n\t")
			} else {
				fmt.Printf("%d\t", node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		fmt.Printf("\n")
		queue = queue[size:]
	}
}

// 利用 []int 生成 *TreeNode
func Ints2Tree(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: ints[0]}
	queue := make([]*TreeNode, 1, 2*n)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// 把 *TreeNode 按照行还原成 []int
func Tree2Ints(root *TreeNode) []int {
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				res = append(res, NULL)
			} else {
				res = append(res, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		queue = queue[size:]
	}
	i := len(res)
	for i > 0 && res[i-1] == NULL {
		i--
	}
	return res[:i]
}

// 二叉树转换成前序遍历切片
func Tree2PreOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := make([]int, 0)
	left := Tree2PreOrder(root.Left)
	right := Tree2PreOrder(root.Right)

	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)
	return res
}

// 二叉树转换成中序遍历切片
func Tree2InOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := make([]int, 0)
	left := Tree2InOrder(root.Left)
	right := Tree2InOrder(root.Right)

	res = append(res, left...)
	res = append(res, root.Val)
	res = append(res, right...)
	return res
}

// 二叉树转换成后序遍历切片
func Tree2PostOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := make([]int, 0)
	left := Tree2PostOrder(root.Left)
	right := Tree2PostOrder(root.Right)

	res = append(res, left...)
	res = append(res, right...)
	res = append(res, root.Val)
	return res
}

//-------------------------------------------
// 把 inorder 和 postorder 切片转换成 二叉树
func InPost2Tree(in, post []int) *TreeNode {
	return nil
}

// 把 preorder 和 inorder 切片转换成 二叉树
func PreIn2Tree(pre, in []int) *TreeNode {
	return nil
}

// 二叉树的最大深度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	return max(left, right) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var flag int

func PrintBinaryTreeGraph(root *TreeNode) {
	height := MaxDepth(root)
	flag = height + 2
	for i := 0; i < height; i++ {
		printRow(root, i)
	}
}

func printRow(root *TreeNode, depth int) {
	placeholder := -1
	vec := []int{}
	getLine(root, depth, &vec)

	setw(flag)
	flag--
	toggle := true
	if len(vec) > 1 {
		for v := 0; v < len(vec); v++ {
			if vec[v] != placeholder {
				if toggle {
					fmt.Printf("/")
				} else {
					fmt.Printf(" \\")
					setw(2 * (depth - 1))
				}
			} else {
				if toggle {
					fmt.Printf(" ")
				} else {
					fmt.Printf("  ")
				}
			}
			toggle = !toggle
		}
		fmt.Printf("\n")
		setw(flag)
		flag--
	}
	toggle = true
	for v := 0; v < len(vec); v++ {
		if vec[v] != -1 {
			if toggle {
				fmt.Printf("%d   ", vec[v])
			} else {
				fmt.Printf("%d", vec[v])
				c := int(math.Pow(2.0, float64(flag)+1) - 1)
				setw(c)
			}
		} else {
			if v == 0 {
				fmt.Printf(" ")
			}
			if toggle {
				fmt.Printf(" ")
				// setw(3)
			}
		}
		toggle = !toggle
	}
	fmt.Printf("\n")
}
func getLine(root *TreeNode, depth int, vals *[]int) {
	placeholder := -1
	if depth <= 0 && root != nil {
		*vals = append(*vals, root.Val)
		return
	}
	if root.Left != nil {
		getLine(root.Left, depth-1, vals)
	} else if depth-1 <= 0 {
		*vals = append(*vals, placeholder)
	}
	if root.Right != nil {
		getLine(root.Right, depth-1, vals)
	} else if depth-1 <= 0 {
		*vals = append(*vals, placeholder)
	}
	return
}
func setw(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf(" ")
	}
}

// printf("fjp %d\n",vec.Size());
// for(auto p=vec.begin();p!=vec.end();p++){
//     printf("%d|",*p);
// }
// printf("\n");
