package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	heights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&heights[i])
	}

	minCosts := make([]int, n)
	for i := 1; i < n; i++ {
		minCost := math.MaxInt32
		for _, comp := range []int{i - 1, i - 2} {
			if comp >= 0 {
				cost := minCosts[comp] + int(math.Abs(float64(heights[i])-float64(heights[comp])))
				if cost < minCost {
					minCost = cost
				}
			}
		}

		minCosts[i] = minCost
	}

	fmt.Println(minCosts[n-1])
}
