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
	var s, t string
	fmt.Fscan(reader, &s, &t)

	m := len(s)
	n := len(t)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			v := 0
			if s[i-1] == t[j-1] {
				v = 1
			}
			dp[i][j] = max(dp[i][j], dp[i-1][j-1]+v)
			dp[i][j] = max(dp[i][j], dp[i-1][j])
			dp[i][j] = max(dp[i][j], dp[i][j-1])
		}
	}

	ansStr := ""
	i := m
	j := n
	for i >= 1 && j >= 1 {
		if s[i-1] == t[j-1] && dp[i][j] == dp[i-1][j-1]+1 {
			ansStr = string(s[i-1]) + ansStr
			i--
			j--
		} else {
			if dp[i-1][j] > dp[i][j-1] {
				i--
			} else {
				j--
			}
		}
	}

	fmt.Println(ansStr)
}
