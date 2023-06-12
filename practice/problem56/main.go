package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)
	nums := make([]int, n+1)
	limits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &nums[i], &limits[i])
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	unableNumber := 100000
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if j == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = unableNumber
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if dp[i-1][j] != unableNumber {
				dp[i][j] = 0
			}

			if j >= nums[i] && dp[i][j-nums[i]] < limits[i] {
				dp[i][j] = dp[i][j-nums[i]] + 1
			}
		}
	}

	if dp[n][m] != unableNumber {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
