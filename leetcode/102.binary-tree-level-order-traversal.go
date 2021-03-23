/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	res := make([][]int, 0)
	for len(queue) > 0 {
		levelQ := make([]*TreeNode, 0)
		levelR := make([]int, 0)
		for _, node := range queue {
			levelR = append(levelR, node.Val)
			left := node.Left
			right := node.Right
			if left != nil {
				levelQ = append(levelQ, left)
			}
			if right != nil {
				levelQ = append(levelQ, right)
			}
		}

		res = append(res, levelR)
		if len(levelQ) > 0 {
			queue = levelQ
		} else {
			break
		}
	}
	return res
}

// @lc code=end

