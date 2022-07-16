/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
package main

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	q := make([]int, 0, k) //单调小队列,存位置下标
	for i := 0; i < len(nums); i++ {
		if len(q) > 0 && q[0] <= i-k { // 判断队头是否出队
			q = q[1:]
		}
		// 维护队列单调性;队头大于队尾,队头出队后才有第二大的
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i) // 下标入队，便于队头出队
		if i >= k-1 {    // 取队头作为窗口最大元素
			res = append(res, nums[q[0]])
		}
	}
	return res
}

// @lc code=end
