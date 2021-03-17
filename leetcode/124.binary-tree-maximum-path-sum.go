/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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
// 最大路径和 = Max {左子树最大路径和，右子树最大路径和， 左子树单边最大路径和 + root.val + 右子树单边最大路径和}
// 单边最大路径和 = Max {左子树单边最大路径和 + root.val, 右子树单边最大路径和 + root.val}

type Result struct {
	singleMax int
	pathMax   int
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := helper(root)
	return res.pathMax
}

func helper(root *TreeNode) Result {
	if root == nil {
		return Result{
			singleMax: 0,
			pathMax:   math.MinInt32, // node.Val可能为负数，所以不能把空节点的最大路径和设为0
		}
	}

	leftRes := helper(root.Left)
	rightRes := helper(root.Right)

	var res Result
	res.singleMax = max2(leftRes.singleMax, rightRes.singleMax) + root.Val
	if res.singleMax < 0 {
		res.singleMax = 0 // 单边最大路径可以不选节点
	}
	res.pathMax = max3(leftRes.pathMax, rightRes.pathMax, leftRes.singleMax+root.Val+rightRes.singleMax)
	return res
}

func max2(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func max3(x, y, z int) int {
	tmp := max2(x, y)
	if tmp > z {
		return tmp
	} else {
		return z
	}
}

// @lc code=end

