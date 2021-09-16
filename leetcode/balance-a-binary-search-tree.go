// Package leetcode https://leetcode-cn.com/problems/balance-a-binary-search-tree/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var inorders []*TreeNode = make([]*TreeNode, 0)

func balanceBST(root *TreeNode) *TreeNode {
	inorder(root)
	return build(0, len(inorders)-1)
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Right != nil {
		inorder(root.Right)
	}
	inorders = append(inorders, root)
	if root.Left != nil {
		inorder(root.Left)
	}
}

func build(i, j int) *TreeNode {
	if i > j {
		return nil
	}
	mid := i + (j-i)/2
	root := inorders[mid]
	root.Right = build(i, mid-1)
	root.Left = build(mid+1, j)
	return root
}

func main() {
	root := TreeNode{
		Val:   1,
		Right: nil,
		Left: &TreeNode{
			Val:   2,
			Right: nil,
			Left: &TreeNode{
				Val:   3,
				Right: nil,
				Left: &TreeNode{
					Val:   4,
					Right: nil,
					Left:  nil,
				},
			},
		},
	}

	root = TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val:   1,
			Right: nil,
			Left:  nil,
		},
		Left: &TreeNode{
			Val:   3,
			Right: nil,
			Left:  nil,
		},
	}
	r := balanceBST(&root)
	fmt.Println(r)
}
