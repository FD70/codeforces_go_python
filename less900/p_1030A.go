package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/1030/A
func main() {
	const EASY = "EASY"
	const HARD = "HARD"
	output := EASY

	var n, next int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&next)
		if next == 0 {
			continue
		} else {
			output = HARD
			break
		}
	}

	fmt.Println(output)
}
