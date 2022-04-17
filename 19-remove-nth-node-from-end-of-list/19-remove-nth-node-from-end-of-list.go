/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    root := head
    
    nodeLen := 0
    for head != nil {
        head = head.Next
        nodeLen += 1
    }
    
    if nodeLen == n {
        return root.Next
    }
    
    head = root
    prev := head
    for i := 1; i < nodeLen; i++ {
        prev = head
        head = head.Next
        if i == nodeLen - n {
            if head != nil {
                prev.Next = head.Next
            } else {
                prev.Next = nil
            }
            return root
        }
    }
    return root
}