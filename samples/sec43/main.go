package main

import "fmt"

func gcd(m, n int) int {
	r := m % n
	if r == 0 {
		return n
	}

	return gcd(n, r)
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(gcd(m, n))
}
