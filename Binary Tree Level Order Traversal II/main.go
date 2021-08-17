package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseArr(arr [][]int) [][]int {
	newArr := make([][]int, len(arr))
	for i, v := range arr {
		newArr[len(arr)-i-1] = v
	}
	return newArr
}

func helper(node *TreeNode, arr [][]int, height int) [][]int {
	if node == nil {
		return arr
	}
	if len(arr) <= height {
		arr = append(arr, []int{})
	}
	arr = helper(node.Left, arr, height+1)
	arr = helper(node.Right, arr, height+1)
	arr[height] = append(arr[height], node.Val)
	return arr
}

func levelOrderBottom(root *TreeNode) [][]int {
	var arr [][]int
	height := 0
	arr = helper(root, arr, height)
	return reverseArr(arr)
}

func main() {
	tree := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	arr := levelOrderBottom(tree)
	fmt.Println(arr)

	arr = levelOrderBottom(nil)
	fmt.Println(arr)
}
