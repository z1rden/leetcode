package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		&ListNode{Val: 7, Next: &ListNode{Val: 2}})

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	l = addTwoNumbersGreat(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		&ListNode{Val: 7, Next: &ListNode{Val: 2}})

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

/*
Суть алгоритма: просто перебор связанных списков.
Временная сложность алгоритма: O(n), где n-большая длина среди двух списков.
Пространственная сложность: O(1).
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	end := head

	var remainder int
	for l1 != nil || l2 != nil {
		var num int

		if l1 == nil {
			num = l2.Val + remainder
			l2 = l2.Next
		} else if l2 == nil {
			num = l1.Val + remainder
			l1 = l1.Next
		} else {
			num = l1.Val + l2.Val + remainder
			l1 = l1.Next
			l2 = l2.Next
		}

		remainder = 0

		if num <= 9 {
			end.Next = &ListNode{Val: num}
		} else {
			remainder = num / 10
			num = num % 10
			end.Next = &ListNode{Val: num, Next: nil}
		}

		end = end.Next
	}

	if remainder > 0 {
		end.Next = &ListNode{Val: remainder, Next: nil}
	}

	return head.Next
}

func addTwoNumbersGreat(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	end := head

	var remainder int
	for l1 != nil || l2 != nil || remainder != 0 {
		num := remainder

		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}

		end.Next = &ListNode{Val: num % 10, Next: nil}
		end = end.Next
		remainder = num / 10

	}

	return head.Next
}
