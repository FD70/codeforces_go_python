package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/71/A
func main() {
	var n int // number of strings

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var next_str string

		fmt.Scan(&next_str)

		len_ns := len(next_str)
		if len_ns <= 10 {
			fmt.Println(next_str)
		} else {
			fmt.Print(next_str[:1], len_ns-2, next_str[len_ns-1:], "\n")
		}
	}

}
