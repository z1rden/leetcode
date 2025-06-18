package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("AB", 1))
}

/*
Суть алгоритма в двойном указателе:
1. Поочередно считываем символы;
2. Второй указатель в это время работает по следующему правилу:
  - если он достиг numrows - 1 (верхней границы записи), то идем в обратную сторону ( 3 2 1 )
  - если он достиг 0 (нижней границы записи), то идем в обратную сторону (1 2 3)

Временная сложность: O(n).
Пространственная сложность: O(1).
*/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([][]byte, numRows)
	j := 0
	d := 1
	for i := 0; i < len(s); i++ {
		rows[j] = append(rows[j], s[i])

		if j == numRows-1 {
			d = -1
		} else if j == 0 {
			d = 1
		}

		j += d
	}

	var result []byte
	for i := 0; i < numRows; i++ {
		result = append(result, rows[i]...)
	}

	return string(result)
}
