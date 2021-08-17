package main

import (
	"fmt"
)

func printTree(node *TreeNode) {
	queue := []*TreeNode{}
	queue = append(queue, node)
	for ; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		if curr != nil {
			if curr.Left != nil || curr.Right != nil {
				queue = append(queue, curr.Left, curr.Right)
			}
			fmt.Print(curr.Val, " ")
		} else {
			fmt.Print("null ")
		}
	}
	fmt.Println()
}

func main() {
	arr := []int{-10, -3, 0, 5, 9}
	list := sliceToList(arr)
	printTree(sortedListToBST(list))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newListNode(val int) *ListNode {
	node := new(ListNode)
	node.Val = val
	return node
}

func newTreeNode(val int) *TreeNode {
	node := new(TreeNode)
	node.Val = val
	return node
}

func sliceToList(arr []int) *ListNode {
	head := new(ListNode)
	node := head
	for _, v := range arr {
		next := newListNode(v)
		node.Next = next
		node = node.Next
	}
	return head.Next
}

func listToSlice(head *ListNode) []int {
	arr := []int{}
	for ; head != nil; head = head.Next {
		arr = append(arr, head.Val)
	}
	return arr
}

func sortedListToBST(head *ListNode) *TreeNode {
	arr := listToSlice(head) 
	var helper func(arr []int) *TreeNode
	helper = func(arr []int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		i := len(arr) / 2
		node := newTreeNode(arr[i])
		node.Left = helper(arr[:i])
		node.Right = helper(arr[i+1:])
		return node
	}
	return helper(arr)
}
