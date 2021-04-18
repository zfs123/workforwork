/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func inorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        result = append(result, root.Val)
        dfs(root.Right)
    }
    dfs(root)
    return result
}

// 非递归
func inorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        top := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        result = append(result, top.Val)
        root = top.Right
    }
    return result
}