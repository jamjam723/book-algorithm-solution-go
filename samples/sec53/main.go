package main

import (
	"fmt"
	"math"
)

func minCostTo(minCosts []int, heights []int, i int) int {
	if minCosts[i] < math.MaxInt32 {
		return minCosts[i]
	}

	minCost := math.MaxInt32
	if i-1 >= 0 {
		minCost = minCostTo(minCosts, heights, i-1) + int(math.Abs(float64(heights[i]-heights[i-1])))
	}

	if i-2 >= 0 {
		tmp := minCostTo(minCosts, heights, i-2) + int(math.Abs(float64(heights[i]-heights[i-2])))
		if tmp < minCost {
			minCost = tmp
		}
	}

	minCosts[i] = minCost
	return minCost
}

func main() {
	var n int
	fmt.Scan(&n)
	heights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&heights[i])
	}

	minCosts := make([]int, n)
	for i := 1; i < n; i++ {
		minCosts[i] = math.MaxInt32
	}
	fmt.Println(minCostTo(minCosts, heights, n-1))
}
