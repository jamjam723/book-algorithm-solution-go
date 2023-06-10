package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&nums[i])
	}

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
			}

			if j >= nums[i] && dp[i-1][j-nums[i]] {
				dp[i][j] = true
			}
		}
	}

	if dp[n][m] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
