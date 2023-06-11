package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// solved the following problem.
// https://onlinejudge.u-aizu.ac.jp/courses/library/7/DPL/1/DPL_1_C
func main() {
	var n, w int
	fmt.Scan(&n, &w)

	values := make([]int, n+1)
	weights := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&values[i])
		fmt.Scan(&weights[i])
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= weights[i] {
				dp[i][j] = max(dp[i][j], dp[i][j-weights[i]]+values[i])
			}
		}
	}

	fmt.Println(dp[n][w])
}
