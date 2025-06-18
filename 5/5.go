package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}

/*
Суть алгоритма:
1. Поочередно считываем элементы строки;
2. Свойство палиндрома: относительно центра он симметричен;
3. Фиксируем позицию и идем в левую и правую стороны от него до тех пор, пока очередные элементы равны (этот случай рассматривается как для четного палиндрома anna, так и для нечетного annna)
4. Из двух длин выбирается максимальная.
5. Сравнивается с максимальной длиной палиндрома всей строки, а не текущей итерации. Если больше, то максимальная длина меняется и считаются новые позиции элементов для возврата.
Временная сложность алгоритма: O(n^2), где n-длина строки s.
Пространственная сложность: O(1).
*/
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	left, right, maxLen := 0, 0, 0
	for i := 0; i < len(s)-1; i++ {
		oddLen := findLenPalindrome(s, i, i+1)
		evenLen := findLenPalindrome(s, i, i)

		curLen := max(oddLen, evenLen)
		if curLen > maxLen {
			maxLen = curLen

			left = i - (maxLen-1)/2
			right = i + maxLen/2
		}
	}

	return s[left : right+1]
}

func findLenPalindrome(s string, left int, right int) int {
	for left >= 0 && right <= len(s)-1 {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}

	return right - left - 1
}
