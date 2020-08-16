/* Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
*/

package main

import (
	"fmt"
	"testing"
)

func Example() {
	tree := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	arr := levelOrderBottom(tree)
	fmt.Println(arr)

	arr = levelOrderBottom(nil)
	fmt.Println(arr)

	// Output:
	// [[15 7] [9 20] [3]]
	// []
}

func Benchmark(b *testing.B) {
	tree := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	for i := 0; i < b.N; i++ {
		levelOrderBottom(tree)
	}
}
