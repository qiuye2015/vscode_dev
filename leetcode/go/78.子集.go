/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
package main

// @lc code=start
func subsets(nums []int) [][]int {
	ret := [][]int{{}}
	for _, num := range nums {
		size := len(ret)
		for i := 0; i < size; i++ {
			tmp := append([]int{}, ret[i]...) // 复制ret[i]
			tmp = append(tmp, num)
			ret = append(ret, tmp)
		}
	}
	return ret
}

// @lc code=end
