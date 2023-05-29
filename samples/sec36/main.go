package main

import "fmt"

func main() {
	var N, W int
	fmt.Scanf("%d %d", &N, &W)

	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}

	var exist bool
	for bit := 0; bit < (1 << N); bit++ {
		partialSum := 0
		for i := 0; i < N; i++ {
			if (bit & (1 << i)) > 0 {
				partialSum += a[i]
			}
		}

		if partialSum == W {
			exist = true
			break
		}
	}

	if exist {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
