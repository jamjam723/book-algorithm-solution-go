package main

import "fmt"

func partialSum(nums []int, cur int, rest int) bool {
	if rest == 0 {
		return true
	}

	if cur < 0 {
		return false
	}

	if partialSum(nums, cur-1, rest) || partialSum(nums, cur-1, rest-nums[cur]) {
		return true
	}

	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	var q int
	fmt.Scan(&q)
	var query int
	for i := 0; i < q; i++ {
		fmt.Scan(&query)
		if partialSum(nums, n-1, query) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
