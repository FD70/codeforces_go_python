package main

import (
	"fmt"
)

func main() {
	var w int

	fmt.Scanf("%d", &w)

	if (w > 3) && (w%2) == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

// https://codeforces.com/problemset/problem/4/A
