package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)
	costs := make([][]int, n+1)
	for i := range costs {
		costs[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			fmt.Fscan(reader, &costs[i][j])
		}
	}

	minTotalCosts := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minTotalCosts[i] = math.MaxInt32
	}

	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			totalCost := minTotalCosts[j] + costs[j][i]
			if totalCost < minTotalCosts[i] {
				minTotalCosts[i] = totalCost
			}
		}
	}

	fmt.Println(minTotalCosts[n])
}
