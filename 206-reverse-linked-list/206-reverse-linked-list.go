/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    next := head.Next
    prev := head
    prev.Next = nil

    head = next
    next = head.Next
    head.Next = prev
    
    for {
        prev := head
        
        if next == nil {
            return head
        }
        head = next
        next = head.Next
        head.Next = prev
    }
}