/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	res := make([][]int, 0)
	fromLeft := true
	for len(queue) > 0 {
		levelQ := make([]*TreeNode, 0)
		levelR := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			var node *TreeNode
			if fromLeft {
				node = queue[i]
			} else {
				node = queue[len(queue)-1-i]
			}
			levelR = append(levelR, node.Val)
		}
		for _, node := range queue {
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
			fromLeft = !fromLeft
		} else {
			break
		}
	}
	return res
}

// @lc code=end

