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
    res, _ := validBstBest(root)
    return res
}

type MinAndMax struct {
    min int
    max int
}
func validBstBest(root *TreeNode) (bool, MinAndMax) {
    var left, right MinAndMax
    var isleft, isright bool

    if root.Left == nil {
        left.min = root.Val
        left.max = root.Val
        isleft = true
    }
    if root.Right == nil {
        right.min = root.Val
        right.max = root.Val
        isright = true
    }

    if root.Left != nil {
        isleft, left = validBstBest(root.Left)
    }
    if root.Right != nil {
        isright, right = validBstBest(root.Right)
    }

    if isleft && isright {
        if root.Left != nil {
            if left.max >= root.Val {
                return false, MinAndMax{}
            }
        }
        if root.Right != nil {
            if right.min <= root.Val {
                return false, MinAndMax{}
            }
        }
        return true, MinAndMax{left.min, right.max}
    } else {
        return false, MinAndMax{}
    }
}