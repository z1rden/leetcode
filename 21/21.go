package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(MergeTwoLists(&ListNode{1, &ListNode{2, nil}}, &ListNode{1, &ListNode{2, nil}}))
}

/*
Суть алгоритма заключается в переборе всех узлов списка.
Изначально необходимо зафиксировать начало (голову) списка, который будет являться конкатенацией двух списков.
Затем начать итерироваться по двум входным спискам, сравнивая их значения.

Сложность алгоритма: O(n+m) - n,m - длины входных цепочек, потому что они перебираются полностью.
Пространственная сложность: O(n+m) n,m - длины входных цепочек, потому что создается новый список длиной n+m.
*/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	begin := &ListNode{}
	end := begin

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			end.Next = l2
			l2 = l2.Next
		} else {
			end.Next = l1
			l1 = l1.Next
		}
		end = end.Next
	}

	if l1 != nil {
		end.Next = l1
	} else {
		end.Next = l2
	}

	return begin.Next
}
