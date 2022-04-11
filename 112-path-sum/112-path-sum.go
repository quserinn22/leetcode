/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    if root.Left == nil && root.Right == nil {
        return targetSum == root.Val
    }
    
    if hasPathSum(root.Left, targetSum - root.Val) {
        return true
    }
    
    if hasPathSum(root.Right, targetSum - root.Val) {
        return true
    }
    
    return false
}