package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindromeWithStr(121))
	fmt.Println(isPalindromeWithInt(121))
}

/*
Сложность алгоритма: O(log10(n)).
В цикле for число x уменьшается в 10 раз на каждой итерации (через x = x / 10), соответственно, количество итераций будет равно log10(n).

Пространственная сложность: O(1).
*/
func isPalindromeWithInt(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0
	xCopy := x

	for x > 0 {
		reverse = (reverse * 10) + (x % 10)
		x = x / 10
	}

	return xCopy == reverse
}

/*
Сложность алгоритма: O(n).
Прогоняется один раз по всему циклу.

Пространственная сложность: O(1).
*/
func isPalindromeWithStr(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	fmt.Println(str)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}
