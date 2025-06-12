package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1}, 1))
	fmt.Println(searchInsert([]int{1}, 2))
	fmt.Println(searchInsert([]int{1}, 0))
}

/*
Суть алгоритма заключается в бинаром поиске.

Сложность алгоритма: O(log(2)n) - n - длина строки.
Принцип "деления пополам": на каждой итерации алгоритм уменьшает зону поиска в 2 раза => отсюда появляется log(2).
Для массива из n элементов максимальное число итераций равно log₂(n).
Пространственная сложность: O(1).
*/
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return left
}
