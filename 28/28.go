package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}

/*
Суть алгоритма заключается в переборе всей входной строки и проверки наличия у очередной подстроки, начинающейся с текущего элемента и оканчивающейся последним элементом слова, второй входной строки.

Сложность алгоритма: O(n*m) - n - длина первой входной строки, m - длина второй входной строки.
То есть выполняется n итераций, где n = len(haystack) +
+ На каждой итерации проверяет, начинается ли haystack[i:] с needle. В худшем случае для каждой позиции i происходит сравнение всех символов needle. Длина needle — m (т.е. m = len(needle)). В худшем случае (например, haystack = "aaa...a", needle = "aaa...b") для каждой позиции i выполняется m сравнений.
Пространственная сложность: O(1).
*/
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if strings.HasPrefix(haystack[i:], needle) {
			return i
		}
	}

	return -1
}
