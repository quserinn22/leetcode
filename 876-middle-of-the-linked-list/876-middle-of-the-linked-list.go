/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    po1 := head
    po2 := head
    for po2 != nil && po2.Next != nil {
        po1 = po1.Next
        po2 = po2.Next
        if po2 != nil {
            po2 = po2.Next
        }
    }
    return po1
}