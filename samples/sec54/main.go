package main

import (
	"fmt"
	"math"
)

func main() {
	var n, w int
	fmt.Scanf("%d %d", &n, &w)

	weights := make([]int, n+1)
	values := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d %d", &weights[i], &values[i])
	}

	knapsack := make([][]int, n+1)
	for i := range knapsack {
		knapsack[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			ifHave := 0
			if j-weights[i] >= 0 {
				ifHave = knapsack[i-1][j-weights[i]] + values[i]
			}
			knapsack[i][j] = int(math.Max(float64(knapsack[i-1][j]), float64(ifHave)))
		}
	}

	fmt.Println(knapsack[n][w])
}
