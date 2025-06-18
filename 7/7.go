package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(120))
}

/*
Временная сложность: O(log10(n)).
Пространственная сложность: O(1).
*/
func reverse(x int) int {
	rev := 0
	for x != 0 {
		rev = rev*10 + x%10
		x = x / 10
	}

	if rev >= math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}

	return rev
}
