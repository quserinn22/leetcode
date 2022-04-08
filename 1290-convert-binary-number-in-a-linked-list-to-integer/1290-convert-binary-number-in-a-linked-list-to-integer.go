/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    ret := head.Val
    head = head.Next
    for head != nil {
        ret *= 2
        ret += head.Val
        head = head.Next
    }
    return ret
}