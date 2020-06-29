package dp

import "fmt"

func Dp(n int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i-1 >= 0 {
				dp[i][j] = dp[i][j] + dp[i-1][j]
			}
			if j-1 >= i {
				dp[i][j] = dp[i][j] + dp[i][j-1]
			}
		}
	}
	return dp[n-1][n-1]
}

func Rob(nums []int) int {
	var dp []int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp = append(dp, nums[0])
		} else if i == 1 {
			if nums[1] > dp[0] {
				dp = append(dp, nums[1])
			} else {
				dp = append(dp, dp[0])
			}
		} else {
			if nums[i]+dp[i-2] > dp[i-1] {
				dp = append(dp, nums[i]+dp[i-2])
			} else {
				dp = append(dp, dp[i-1])
			}
		}
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}

var dps []int

func NumSquares(n int) int {
	dps=make([]int,n+1)
	min:= numItemSquares(1, n)
	fmt.Println(dps)
	return min
}

func numItemSquares(i, n int) int {

	if n == 0 {
		dps[n] = 0
		return 0
	}
	if i*i == n {
		dps[n] = 1
		return 1
	}
	if i*i > n {
		return 2 << 32
	}
	min := 2 << 32
	for j := i; j*j <= n; j++ {
		m := 0
		if dps[n-j*j] != 0 {
			fmt.Println("---ok---")
			m = dps[n-j*j] + 1
		} else {
			m = numItemSquares(j, n-j*j) + 1
		}
		if m <= min {
			min = m
		}
	}
	dps[n] = min
	return min
}
