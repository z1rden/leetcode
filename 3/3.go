package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstringWindow("pwwkew"))
}

// Плохой вариант, потому что многократно обрабатываются одни и те же символы.
func lengthOfLongestSubstring(s string) int {
	var (
		curMap         = make(map[byte]int)
		curLen, maxLen int
	)

	for i := 0; i < len(s); i++ {
		if _, ok := curMap[s[i]]; ok {
			i = curMap[s[i]]
			curLen = 0
			curMap = make(map[byte]int)
		} else {
			curMap[s[i]] = i
			curLen++

			if maxLen < curLen {
				maxLen = curLen
			}
		}
	}

	return maxLen
}

/*
Суть алгоритма в движущемся окне (встретили повтор, двигаем левую границу; не встретили - правую).
Временная сложность алгоритма: O(n), где n-длина строки s.
Пространственная сложность: O(1).
*/
func lengthOfLongestSubstringWindow(s string) int {
	left := 0
	maxLen := 0
	curMap := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		for curMap[s[right]] > 0 {
			delete(curMap, s[left])
			left++

		}

		curMap[s[right]]++
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
