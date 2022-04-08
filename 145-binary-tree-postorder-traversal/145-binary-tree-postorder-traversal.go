/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    ret := []int{}
    ret = append(ret, postorderTraversal(root.Left)...)
    ret = append(ret, postorderTraversal(root.Right)...)
    ret = append(ret, root.Val)
    return ret
}