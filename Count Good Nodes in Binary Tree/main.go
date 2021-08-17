package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{2, nil,
		&TreeNode{4, &TreeNode{10, nil, nil},
			&TreeNode{8, &TreeNode{4, nil, nil}, nil}}}
	fmt.Println(goodNodes(root))
}

func goodNodes(root *TreeNode) int {
	numGoodNodes := 0
	var helper func(*TreeNode, int)
	helper = func(tn *TreeNode, highest int) {
		if tn == nil {
			return
		}
		if tn.Val >= highest {
			numGoodNodes += 1
			highest = tn.Val
		}
		helper(tn.Left, highest)
		helper(tn.Right, highest)
	}
	helper(root, -1e5)
	return numGoodNodes
}
