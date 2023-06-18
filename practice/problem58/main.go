package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)

	heights := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &heights[i])
	}

	means := make([][]float64, n+1)
	for i := range means {
		means[i] = make([]float64, n+1)
	}

	for i := 0; i < n; i++ {
		total := float64(0.0)
		for j := i + 1; j <= n; j++ {
			total += float64(heights[j])
			means[i][j] = total / float64(j-i)
		}
	}

	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = -10000000000.0
		}
	}

	dp[0][0] = 0.0
	for i := 1; i <= n; i++ {
		for j := 1; j <= minInt(i, m); j++ {
			for k := 0; k <= i-1; k++ {
				dp[i][j] = math.Max(dp[i][j], dp[k][j-1]+means[k][i])
			}
		}
	}

	fmt.Println(dp[n][m])
}
