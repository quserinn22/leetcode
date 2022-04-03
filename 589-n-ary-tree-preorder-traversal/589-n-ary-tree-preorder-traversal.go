/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    if root == nil {
        return []int{}
    }
    
    ret := []int{root.Val}
    for _, child := range root.Children {
        ret = append(ret, preorder(child)...)
    }
    return ret
}