/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return head
    }
    ret := head
    
    for head != nil {
        node := head.Next
        for node != nil && node.Val == val {
            node = node.Next
        }
        
        head.Next = node
        head = head.Next
    }
    
    if ret.Val == val {
        return ret.Next
    }
    
    return ret
}