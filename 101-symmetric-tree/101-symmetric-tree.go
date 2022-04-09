/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    
    lo := levelOrder(root)
    fmt.Println(lo)

    for i := 1; i < len(lo); i++ {
        if len(lo[i]) % 2 != 0 {
            return false
        }
        
        for j := 0; j < len(lo[i]) / 2; j++ {
            if lo[i][j] != lo[i][len(lo[i]) - j - 1] {
                return false
            }
        }
    }
    return true
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{{-101}}
    }
    
    ret := [][]int{{root.Val}}
    
    left := levelOrder(root.Left)
    right := levelOrder(root.Right)
    
    
    for i := 0; i < max(len(left), len(right)); i++ {
        if i >= len(left) {
            ret = append(ret, right[i:]...)
            break
        } else if i >= len(right) {
            ret = append(ret, left[i:]...)
            break
        } else {
            ret = append(ret, append(left[i], right[i]...))
        }
    }
    
    return ret
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}