package main

import (
	"fmt"
	"incq.pro/golang-algo/common"
)

func main() {
	root := common.TreeNode{
		Val: 0,
	}
	root.Left = &common.TreeNode{
		Val: 1,
	}
	root.Right = &common.TreeNode{
		Val: 2,
	}
	fmt.Println(binaryTreeLevel(&root))
}

func binaryTreeLevel(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	left := binaryTreeLevel(root.Left)
	right := binaryTreeLevel(root.Right)
	max := left
	if right > max {
		max = right
	}
	return max + 1
}