package main

import "fmt"

func main() {
	input := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	output := hasPathSum(input, 22)
	fmt.Println(output)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	result := false

	var fn func(*TreeNode, int, int)
	fn = func(node *TreeNode, target, prev int) {
		if node.Left != nil {
			fn(node.Left, target, prev+node.Val)
		}
		if node.Right != nil {
			fn(node.Right, target, prev+node.Val)
		}
		if prev+node.Val == target && node.Left == nil && node.Right == nil {
			result = true
		}
	}

	if root != nil {
		fn(root, targetSum, 0)
	}

	return result
}
