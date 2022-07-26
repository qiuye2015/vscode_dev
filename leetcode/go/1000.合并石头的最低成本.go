/*
 * @lc app=leetcode.cn id=1000 lang=golang
 *
 * [1000] 合并石头的最低成本
 */

package main

// @lc code=start

/*
   Check whether we will be able to merge n piles into 1 pile .

       In step-1 we merge k pile and then we are left with n-k+1 piles or n-(k-1);
       In Step-2 we again merge k pile and then we are left with ((n-k+1)-k)+1 or n-2*(k-1);
       In Step-3 we gain merge k pile and left with (((n-k+1)-k+1)-k)+1 or n-3*(k-1)
       .......
       .......
       .......
       After some m steps we should be left with 1 pile
       Therefore , n-m*(k-1) == 1
              (n-1)-m*(k-1)=0;
              Since m needs to be an integer therefore ,
              if (n-1)%(k-1) == 0 ,
              then we can merge otherwise not possible.
*/

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	const N = 50
	const INT_MAX = 99999999999
	dp := [N][N][N]int{}
	prefix := make([]int, n+1)
	sum := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	// Calculating the prefix sum so that sum of any segment[i..j] can be calculated easily
	for i := 0; i < n; i++ {
		sum += stones[i]
		prefix[i+1] = sum // 前缀和
	}
	var mincost func(start, end, piles int) int // pile 堆
	mincost = func(start, end, piles int) int {
		// Cost of converting segment [i..i] into 1 pile is zero
		if start == end && piles == 1 {
			return 0
		}
		// Cost of converting segment[i..i] into other than 1 pile is not possible , so placed MAX value
		if start == end {
			return INT_MAX
		}
		// Check whether the subproblem has already been solved .
		if dp[start][end][piles] != -1 {
			return dp[start][end][piles]
		}
		res := 0
		if piles == 1 { // If segment[i...j] is to converted into 1 pile
			// Here dp(i,j,1) = dp(i,j,k) + sum[i...j]
			res = mincost(start, end, k) + prefix[end+1] - prefix[start]
		} else {
			// dp(i,j,piles) = min( dp(i,j,piles), dp(i,t,1) + dp(t+1,j,piles-1)) for all t E i<=t<j
			cost := INT_MAX
			for t := start; t < end; t++ {
				left := mincost(start, t, 1)
				right := mincost(t+1, end, piles-1)
				cost = min(cost, left+right)
			}
			res = cost
		}
		dp[start][end][piles] = res
		return res
	}
	return mincost(0, n-1, 1)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
