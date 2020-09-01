/*
Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var l3 *ListNode
	if l1.Val > l2.Val {
		l3 = l2
		l2 = l2.Next
	} else {
		l3 = l1
		l1 = l1.Next
	}
	head := l3
	for l1 != nil && l2 != nil {
		if l2.Val >= l1.Val {
			l3.Next = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
		}
		l3 = l3.Next
	}
	if l1 != nil {
		l3.Next = l1
	} else {
		l3.Next = l2
	}

	return head
}
