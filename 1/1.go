package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumFastest([]int{2, 7, 11, 15}, 9))
}

/*
Сложность алгоритма: O(n^2).
Для i = 0: выполняется n-1 сравнений.
Для i = 1: выполняется n-2 сравнений.
...
Для i = n-1: выполняется 0 сравнений.
То есть высокая сложность будет обусловлена использованием вложенных циклов.
Просуммировав эти позиции, получается n*(n-1)/2. Это выражение при n->inf стремится к n^2, что и определяет сложность алгоритма.
Пространственная сложность: O(1).
Алгоритм не использует дополнительные структуры данных - только константные данные:
1. i,j
2. []int{i,j}
*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}

	return nil
}

/*
Сложность алгоритма: O(n).
Алгоритм использует один прогон по массиву.
Пространственная сложность: O(n).
Это получается из-за использования дополнительной хэш-таблицы values.

Таким образом, жертвуется память O(1)->O(n) для ускорения времени O(n^2)->O(n).
*/
func twoSumFastest(nums []int, target int) []int {
	values := map[int]int{}

	for i := 0; i < len(nums); i++ {
		current := nums[i]
		x := target - current

		if _, exists := values[x]; exists {
			return []int{values[x], i}
		}

		values[current] = i
	}

	return nil
}
