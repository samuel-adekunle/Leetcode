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


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	carry := 0
	helper(res, l1, l2, carry)
	return res
}
