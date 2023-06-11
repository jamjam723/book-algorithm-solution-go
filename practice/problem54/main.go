package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// solved the following problem.
// https://algo-method.com/tasks/312
func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&nums[i])
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if j == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = 505
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i] {
				dp[i][j] = min(dp[i][j], dp[i-1][j-nums[i]]+1)
			}
		}
	}

	if dp[n][m] <= k {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
