package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

/*
Суть алгоритма заключается в переборе всех чисел.
Изначально необходимо зафиксировать первый элемент среза и его позицию.
Если следующий элемент не равен зафиксированному, то на следующее место от текущего ставится новый элемент.
Такая операция проделывается со всем списком.

Сложность алгоритма: O(n) - n - длина входного среза.
Пространственная сложность: O(1).
*/
func removeDuplicates(nums []int) int {
	nextPlace := 0
	current := nums[nextPlace]
	nextPlace++
	for i := 1; i < len(nums); i++ {
		if current != nums[i] {
			nums[nextPlace] = nums[i]
			nextPlace++
			current = nums[i]
		}
	}

	return nextPlace
}
