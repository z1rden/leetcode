package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefixBinarySearch([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefixDAndC([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefixVScan([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefixLCP([]string{"flower", "flow", "flight"}))
}

/*
Суть алгоритма заключается в методе бинарного поиска.
Для начала необходимо найти слово с минимальной длиной. После у этого слова согласно методу бинарного поиска берется его часть и проверяется, присутствует ли эта часть во всех других словах.
Если присутствует, то добавляется к этому префиксу еще букв. Если нет, то количество букв уменьшается.
Сложность алгоритма: O(S*logm), S-сумма всех символов в строках, где m - количество буков в одном из n одинаковых строк.
Алгоритм выполняет logm итераций, для каждой из которых требуется S=m⋅n сравнений
Пространственная сложность: O(1).
*/
func longestCommonPrefixBinarySearch(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := math.MaxInt
	for i := 0; i < len(strs); i++ {
		minLen = min(minLen, len(strs[i]))
	}

	low, high := 0, minLen
loop:
	for low <= high {
		mid := (low + high) / 2

		str := strs[0][:mid]
		for i := 1; i < len(strs); i++ {
			if !strings.HasPrefix(strs[i], str) {
				high = mid - 1

				continue loop
			}
		}
		low = mid + 1
	}

	return strs[0][:(low+high)/2]
}

/*
Суть алгоритма заключается в ассоциативном свойстве LCP, LCP(S1...Sn)=LCP(LCP(S1...Si),LCP(Si+1...Sn).
Для этого задача поиска LCP(Si...Sj) разбивается на две подзадачи LCP(Si...Smid) и LCP(Smid+1...Sj), где mid = (i+j)/2.
Пусть их решения lcpLeft и lcpRight. Они сраниваются до тех пор, пока не найдется совпадение, кое является LCP(Si...Sj).

Сложность алгоритма: O(S), S-сумма всех символов в строках.
В худшем случае будет n одинаковых строк длиной m, и алгоритм выполнит S=m⋅n сравнений символов.
Несмотря на то, что наихудший случай остаётся таким же, как в подходе ниже, в лучшем случае выполняется не более n⋅minLen сравнений, где minLen — длина самой короткой строки в массиве.
Пространственная сложность: O(m*log10n).
Поскольку мы сохраняем рекурсивные вызовы в стеке выполнения, требуется дополнительная память. Существует logn рекурсивных вызовов, для каждого из которых требуется m единиц памяти для хранения результата, поэтому временная сложность составляет O(m⋅logn)
DAndC - divide & conquere.
*/
func longestCommonPrefixDAndC(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	LCP := func(str1, str2 string) string {
		minLen := min(len(str1), len(str2))
		for i := 0; i < minLen; i++ {
			if str1[i] != str2[i] {
				return str1[:i]
			}
		}

		return str1[:minLen]
	}

	var divideAndConquer func([]string, int, int) string
	divideAndConquer = func(strs []string, left int, right int) string {
		if left == right {
			return strs[left]
		} else {
			mid := (left + right) / 2
			lcpLeft := divideAndConquer(strs, left, mid)
			lcpRight := divideAndConquer(strs, mid+1, right)

			return LCP(lcpLeft, lcpRight)
		}
	}

	return divideAndConquer(strs, 0, len(strs)-1)
}

/*
Суть алгоритма заключается в сравнении всех строк между собой по вертикали, то есть сначала проверяется первый символ у всех, затем второй и тд.
Сложность алгоритма: O(S), S-сумма всех символов в строках.
В худшем случае будет n одинаковых строк длиной m, и алгоритм выполнит S=m⋅n сравнений символов.
Несмотря на то, что наихудший случай остаётся таким же, как в подходе ниже, в лучшем случае выполняется не более n⋅minLen сравнений, где minLen — длина самой короткой строки в массиве.
Пространственная сложность: O(1).
*/
func longestCommonPrefixVScan(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var prefix string
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return prefix
			}
		}
		prefix += string(c)
	}

	return prefix
}

/*
Можно заметить, что LCP(S1,S2,...,Sn)=LCP(LCP(...LCP(S1,S2)...)Sn).
LCP (largest contentful paint) - это в данном случае наибольший префикс между двумя словами.
Когда LCP(S1,Si) становится пустой строкой, алгоритм завершается.
Сложность алгоритма: O(S), S-сумма всех символов в строках.
Пространственная сложность: O(1).
*/
func longestCommonPrefixLCP(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	lcp := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], lcp) {
			lcp = lcp[:len(lcp)-1]
			if lcp == "" {
				return ""
			}
		}
	}

	return lcp
}
