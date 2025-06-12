package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord("Today is a nice day"))
}

/*
Суть алгоритма:
1. Убираются все пробелы с конца до непробельного символа. Этот символ запоминается.
2. Затем с этого символа начинается справа-налево поиск первого пробела.
3. Если таковой был найден до начала, то вернуть разницу позиций запомненного символа и найденного, иначе их разницу + 1.

Сложность алгоритма: O(n) - n - длина строки.
Пространственная сложность: O(1).
*/
func lengthOfLastWord(s string) int {
	end := len(s) - 1
	for s[end] == ' ' {
		end--
	}

	begin := end
	for s[begin] != ' ' {
		if begin == 0 {
			return end - begin + 1
		}
		begin--
	}

	return end - begin
}
