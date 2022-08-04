/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

package main

// @lc code=start
func sortColors(nums []int) {
	const N = 3
	vec := [N]int{}
	for _, num := range nums {
		vec[num]++
	}
	k := 0
	for i := 0; i < N; i++ {
		for j := 0; j < vec[i]; j++ {
			nums[k] = i
			k++
		}
	}
}

// func sortColors(nums []int)  {
//     zero:=-1 //0的右边界      [0,   zero]
//     two:=len(nums)//2的左边界 [two, len(nums)-1]
//     for i:=0;i<two;{
//         if nums[i]==0{
//             zero++
//             nums[i],nums[zero]=nums[zero],nums[i]
//             i++
//         }else if nums[i]==2 {
//             two--
//             nums[i],nums[two]=nums[two],nums[i]
//         }else if nums[i]==1{
//             i++
//         }
//     }
// }
// @lc code=end
