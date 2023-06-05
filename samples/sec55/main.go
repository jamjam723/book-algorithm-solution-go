package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	var s string
	fmt.Scan(&s)
	var t string
	fmt.Scan(&t)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, row := range dp {
		for j := range row {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[0][0] = 0
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i > 0 && j > 0 {
				editCost := 1
				if s[i-1] == t[j-1] {
					editCost = 0
				}
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+editCost)
			}

			if i > 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
			}
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
			}
		}
	}
	fmt.Println(dp[m][n])
}
