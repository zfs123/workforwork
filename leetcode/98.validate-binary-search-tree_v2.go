/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	res := make([]int, 0)
	visit(root, &res)
	if len(res) == 1 {
		return true
	}
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

func visit(root *TreeNode, queue *[]int) {
	if root == nil {
		return
	}

	visit(root.Left, queue)
	*queue = append(*queue, root.Val)
	visit(root.Right, queue)
}