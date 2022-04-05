/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    po1 := head
    po2 := head
    
    for {
        po1 = po1.Next
        if po1 == nil {
            return false
        }
        
        po2 = po2.Next
        if po2 == nil {
            return false
        }
        po2 = po2.Next
        if po2 == nil {
            return false
        }
        
        if po1 == po2 || po1 == po2.Next {
            return true
        }
    }
}