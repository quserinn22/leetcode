/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    read := head
    node := head
    cache := map[int]bool{
        head.Val: true,
    }
    for {
        if read == nil {
            node.Next = nil
            return head
        }
        
        if _, ok := cache[read.Val]; ok {
            read = read.Next
            continue
        } else {
            cache[read.Val] = true
            node.Next = read
            node = node.Next
            read = read.Next
        }
    }
}