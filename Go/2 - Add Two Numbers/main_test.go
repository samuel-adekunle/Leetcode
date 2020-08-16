/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

package main

import (
	"fmt"
	"testing"
)

func Example() {
	n1 := ListNode{2, &ListNode{4, &ListNode{9, nil}}}
	n2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	res := addTwoNumbers(&n1, &n2)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}

	// Output:
	// 7041
}

func Benchmark(b *testing.B) {
	n1 := ListNode{2, &ListNode{4, &ListNode{9, nil}}}
	n2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	for i := 0; i < b.N; i++ {
		addTwoNumbers(&n1, &n2)
	}
}
