package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToIntPerfect("MCMXCIV"))
}

/*
Суть алгоритма заключается в том, что все числа загоняются в мапу, а числа смотрятся как бы попарно, чтобы отметить интересующие пары:
I можно поставить перед V (5) и X (10), чтобы получить 4 и 9.
X можно поставить перед L (50) и C (100), чтобы получить 40 и 90.
C можно поставить перед D (500) и M (1000), чтобы получить 400 и 900.
Можно заметить, что если следующее число за текущим его меньше, то текущее необходимо вычесть, а следующее прибавлять как и обычно (к примеру, IV, -1+5=4).

Сложность алгоритма O(n), так как один проход по всему слову.
Пространственная сложность O(1), так как вводятся только константы, не зависящие от входных данных.
*/

func romanToIntPerfect(s string) int {
	var sum int
	dict := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s)-1; i++ {
		if dict[string(s[i])] < dict[string(s[i+1])] {
			sum -= dict[string(s[i])]
		} else {
			sum += dict[string(s[i])]
		}
	}

	return sum + dict[string(s[len(s)-1])]
}

/*
Некрасивый вариант.
*/

func romanToInt(s string) int {
	var sum int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i != len(s)-1 {
				if s[i+1] == 'V' {
					sum += 4
					i++
					continue
				} else if s[i+1] == 'X' {
					sum += 9
					i++
					continue
				}
			}
			sum += 1
		case 'V':
			sum += 5
		case 'X':
			if i != len(s)-1 {
				if s[i+1] == 'L' {
					sum += 40
					i++
					continue
				} else if s[i+1] == 'C' {
					sum += 90
					i++
					continue
				}
			}
			sum += 10
		case 'L':
			sum += 50
		case 'C':
			if i != len(s)-1 {
				if s[i+1] == 'D' {
					sum += 400
					i++
					continue
				} else if s[i+1] == 'M' {
					sum += 900
					i++
					continue
				}
			}
			sum += 100
		case 'D':
			sum += 500
		case 'M':
			sum += 1000
		}
	}

	return sum
}
