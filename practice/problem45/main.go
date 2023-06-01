package main

import (
	"fmt"
	"strconv"
)

func countShitigosan(nums string, maxNum int, used int) int {
	count := 0

	if used == 0b111 {
		count += 1
	}

	for i, v := range []string{"3", "5", "7"} {
		candidate, _ := strconv.Atoi(nums + v)
		if candidate <= maxNum {
			count += countShitigosan(nums+v, maxNum, used|1<<i)
		}
	}

	return count
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(countShitigosan("", n, 0b000))
}
