/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    solve(root)
    return root
}

func solve(head *Node) {
    if head == nil {
        return
    }
    
    if head.Left != nil {
        head.Left.Next = head.Right
    }
    
    if head.Right != nil {
        if head.Next != nil {
            head.Right.Next = head.Next.Left
        }
    }
    
    solve(head.Left)
    solve(head.Right)
}