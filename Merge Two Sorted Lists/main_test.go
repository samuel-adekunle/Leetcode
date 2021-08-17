/*
Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4 */

package main

import (
	"fmt"
	"testing"
)

func Example() {
	l1, l2 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l3 := mergeTwoLists(l1, l2)
	for l3 != nil {
		fmt.Print(l3.Val, " ")
		l3 = l3.Next
	}
	l3 = mergeTwoLists(&ListNode{2, nil}, &ListNode{1, nil})
	for l3 != nil {
		fmt.Print(l3.Val, " ")
		l3 = l3.Next
	}
	//Output:
	// 1 1 2 3 4 4 1 2
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l1, l2 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
		mergeTwoLists(l1, l2)
	}
}
