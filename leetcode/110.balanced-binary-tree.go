/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, res := helper(root)
	return res
}

func helper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, leftBalanced := helper(root.Left)
	rightDepth, rightBalanced := helper(root.Right)
	if leftBalanced == false || rightBalanced == false {
		return 0, false
	}

	if leftDepth > rightDepth {
		if leftDepth-rightDepth > 1 {
			return 1 + leftDepth, false
		} else {
			return 1 + leftDepth, true
		}
	} else {
		if rightDepth-leftDepth > 1 {
			return 1 + rightDepth, false
		} else {
			return 1 + rightDepth, true
		}
	}

}

// @lc code=end

