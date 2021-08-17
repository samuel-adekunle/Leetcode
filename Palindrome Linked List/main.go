package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(arr []int) (root *ListNode) {
	head := new(ListNode)
	root = head
	for _, v := range arr {
		next := new(ListNode)
		next.Val = v
		head.Next = next
		head = next
	}
	return root.Next
}

func Length(list *ListNode) (length int) {
	for ; list != nil; list = list.Next {
		length++
	}
	return
}

func Reverse(list *ListNode) *ListNode {
	var prev *ListNode
	for list != nil {
		next := list.Next
		list.Next = prev
		prev = list
		list = next
	}
	return prev
}

func isEqual(a, b *ListNode) bool {
	la := Length(a)
	lb := Length(b)
	if la != lb {
		return false
	}
	for i := 0; i < la; i++ {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return true
}

func isPalindrome(head *ListNode) bool {
	length := Length(head)
	mid := head
	for i := 0; i < (length)/2; i++ {
		mid = mid.Next
	}
	if length > 0 && length%2 == 0 {
		next := mid.Next
		dummy := New([]int{mid.Val})
		dummy.Next = next
		mid.Next = dummy
	}
	mid = Reverse(mid)
	return isEqual(head, mid)
}

func main() {
	arr := []int{-1, -1}
	head := New(arr)
	fmt.Printf("is %v a palindrome: %v\n", arr, isPalindrome(head))
}
