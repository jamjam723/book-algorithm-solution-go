package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	points := make([]int, n+1)
	totalPoint := 0
	var tmp int
	for i := 1; i <= n; i++ {
		fmt.Scan(&tmp)
		points[i] = tmp
		totalPoint += tmp
	}

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, totalPoint+1)
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= totalPoint; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
			}

			if j >= points[i] && dp[i-1][j-points[i]] {
				dp[i][j] = true
			}
		}
	}

	ifAllUsed := dp[n]
	count := 0
	for j := 0; j <= totalPoint; j++ {
		if ifAllUsed[j] {
			count += 1
		}
	}
	fmt.Println(count)
}
