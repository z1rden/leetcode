package main

import "fmt"

func main() {
	fmt.Println(isValidWithBytes("([])"))
	fmt.Println(isValidWithStr("([])"))
}

/*
Временная сложность: O(n). Каждый символ обрабатывается один раз.
Пространственная сложность: O(n), так как в худшем случае стек может хранить все символы (например, все открывающие скобки).
*/
func isValidWithBytes(s string) bool {
	pairsBrackets := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	var stack []byte
	for i := 0; i < len(s); i++ {
		if _, ok := pairsBrackets[s[i]]; !ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairsBrackets[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

/*
Временная сложность: O(n^2).
Каждый символ обрабатывается один раз.
Операция workstr += string(s[i]) в худшем случае (когда все символы — открывающие скобки) имеет сложность:
На первой итерации: копирование 1 символа.
На второй: копирование 2 символов.
...
На последней: копирование n символов.
Суммарно: 1 + 2 + ... + n = n(n+1)/2 → O(n²).

Пространственная сложность: O(n), так как в худшем случае стек может хранить все символы (например, все открывающие скобки).
*/
func isValidWithStr(s string) bool {
	pairsBrackets := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	var workstr string
	for i := 0; i < len(s); i++ {
		if _, ok := pairsBrackets[s[i]]; !ok {
			workstr += string(s[i])
		} else {
			ok, workstr = checkLastChar(workstr, pairsBrackets[s[i]])
			if !ok {
				return false
			}
		}
	}

	if workstr != "" {
		return false
	}

	return true
}

func checkLastChar(str string, c uint8) (bool, string) {
	if len(str) == 0 {
		return false, ""
	}
	if str[len(str)-1] != c {
		return false, ""
	}

	str = str[:len(str)-1]

	return true, str
}
