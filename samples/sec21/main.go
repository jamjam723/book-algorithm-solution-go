package main

import "fmt"

func main() {
	var N, t int
	fmt.Scanf("%d %d", &N, &t)

	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}

	exist := false
	for _, v := range a {
		if v == t {
			exist = true
			break
		}
	}
	fmt.Println(exist)
}
