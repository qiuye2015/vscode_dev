/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

package main

// @lc code=start
// func multiply(num1 string, num2 string) string {
// 	if num1 == "0" || num2 == "0" {
// 		return "0"
// 	}
// 	res := []byte{}
// 	for i := len(num1) - 1; i >= 0; i-- {
// 		x := num1[i] - '0'
// 		var carry byte
// 		tmp := []byte{}
// 		for k := 0; k < len(num1)-1-i; k++ {
// 			tmp = append(tmp, '0')
// 		}
// 		for j := len(num2) - 1; j >= 0 || carry > 0; j-- {
// 			if j >= 0 {
// 				y := num2[j] - '0'
// 				carry = x*y + carry
// 			}
// 			tmp = append(tmp, carry%10+'0')
// 			carry = carry / 10
// 		}
// 		res = addBytes(res, tmp)
// 	}
// 	ans := reserverBytes(res)
// 	return string(ans)
// }
// func reserverBytes(str []byte) []byte {
// 	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
// 		str[i], str[j] = str[j], str[i]
// 	}
// 	return str
// }
// func addBytes(str1, str2 []byte) []byte {
// 	sum := []byte{}
// 	var carry byte
// 	for i, j := 0, 0; i < len(str1) || j < len(str2) || carry > 0; i, j = i+1, j+1 {
// 		if i < len(str1) {
// 			carry += str1[i] - '0'
// 		}
// 		if j < len(str2) {
// 			carry += str2[j] - '0'
// 		}
// 		sum = append(sum, carry%10+'0')
// 		carry = carry / 10
// 	}
// 	return sum
// }
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	ret := make([]byte, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			val := (num1[i] - '0') * (num2[j] - '0')
			ret[i+j+1] += val
			if ret[i+j+1] >= 10 {
				ret[i+j] += ret[i+j+1] / 10
				ret[i+j+1] = ret[i+j+1] % 10
			}
		}
	}
	if ret[0] == 0 {
		ret = ret[1:]
	}
	for i := range ret {
		ret[i] += '0'
	}
	return string(ret)
}

// @lc code=end
