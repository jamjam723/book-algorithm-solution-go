package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)
	hapiness := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		var a, b, c int
		fmt.Fscan(reader, &a, &b, &c)
		hapiness[i] = []int{a, b, c}
	}

	dp := make([][]int, n+1)
	for i := range hapiness {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = 0
	dp[0][1] = 0
	dp[0][2] = 0

	for i := 1; i <= n; i++ {
		dp[i][0] = max(dp[i-1][1]+hapiness[i][0], dp[i-1][2]+hapiness[i][0])
		dp[i][1] = max(dp[i-1][0]+hapiness[i][1], dp[i-1][2]+hapiness[i][1])
		dp[i][2] = max(dp[i-1][0]+hapiness[i][2], dp[i-1][1]+hapiness[i][2])
	}

	tmp := max(dp[n][0], dp[n][1])
	fmt.Println(max(tmp, dp[n][2]))
}
