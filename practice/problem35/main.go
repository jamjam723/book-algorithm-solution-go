package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	count := 0
L:
	for {
		for i := 0; i < n; i++ {
			if nums[i]%2 != 0 {
				break L
			}

			nums[i] /= 2
		}
		count++
	}

	fmt.Println(count)
}
