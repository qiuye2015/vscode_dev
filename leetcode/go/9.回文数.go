/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

package main

// @lc code=start
func isPalindrome(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0， 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reservedNum := 0
	for x > reservedNum {
		reservedNum = reservedNum*10 + x%10
		x /= 10
	}
	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 当输入为 12321 时，循环的末尾 x = 12，revertedNumber = 123，
	return x == reservedNum || x == reservedNum/10
}

// @lc code=end
