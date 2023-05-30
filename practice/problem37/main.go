package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	total := 0
	for bit := 0; bit < (1 << (len(s) - 1)); bit++ {
		head := 0
		for i := 0; i < len(s)-1; i++ {
			if (bit & (1 << i)) > 0 {
				partialNum, _ := strconv.Atoi(s[head : i+1])
				total += partialNum

				head = i + 1
			}
		}

		rest, _ := strconv.Atoi(s[head:])
		total += rest
	}
	fmt.Println(total)
}
