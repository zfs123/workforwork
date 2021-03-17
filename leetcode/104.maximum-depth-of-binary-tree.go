/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1 + Max(maxDepth(root.Left), maxDepth(root.Right))
	return depth
}

func Max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

// @lc code=end

