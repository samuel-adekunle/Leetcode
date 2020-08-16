/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func add(v1, v2, carry int) (v int, c int) {
	val := v1 + v2 + carry
	return val % 10, val / 10
}

func advance(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	}
	return l
}

func getVal(l *ListNode) int {
	if l != nil {
		return l.Val
	}
	return 0
}

func helper(res, l1, l2 *ListNode, carry int) {
	res.Val, carry = add(getVal(l1), getVal(l2), carry)
	l1 = advance(l1)
	l2 = advance(l2)
	if l1 == nil && l2 == nil {
		if carry != 0 {
			res.Next = &ListNode{
				carry, nil,
			}
		}
		return
	}
	res.Next = new(ListNode)
	res = advance(res)
	helper(res, l1, l2, carry)
}

// recursive solution
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	carry := 0
	helper(res, l1, l2, carry)
	return res
}
