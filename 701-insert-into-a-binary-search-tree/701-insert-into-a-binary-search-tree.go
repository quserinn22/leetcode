/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{
            Val: val,
        }
    }
    
    node := root
    prev := root
    for node != nil {
        prev = node
        if node.Val > val {
            node = node.Left
        } else {
            node = node.Right
        }
    }
    
    if prev.Val > val {
        prev.Left = &TreeNode{
            Val: val,
        }
    } else {
        prev.Right = &TreeNode{
            Val: val,
        }
    }
    
    return root
    
}