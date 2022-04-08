/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    ret := 0
    if root.Left != nil {
        if root.Left.Left == nil && root.Left.Right == nil {
            ret += root.Left.Val
        } else {
            ret += sumOfLeftLeaves(root.Left)
        }
    }
    
    if root.Right != nil {
        ret += sumOfLeftLeaves(root.Right)
    }
    
    return ret
}