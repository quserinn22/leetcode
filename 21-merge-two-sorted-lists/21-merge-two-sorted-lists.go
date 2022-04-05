/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    po1 := list1
    po2 := list2
    
    if list1 == nil {
        return list2
    } else if list2 == nil {
        return list1
    }
    
    head := &ListNode{
        Val: min(po1.Val, po2.Val),
    }
    ret := head
    
    for po1 != nil || po2 != nil {
        if po1 == nil {
            node := ListNode{
                Val: po2.Val,
                Next: po2.Next,
            }
            head.Next = &node
            return ret.Next
        } else if po2 == nil {
            node := ListNode{
                Val: po1.Val,
                Next: po1.Next,
            }
            head.Next = &node
            return ret.Next
        } else {
            if po1.Val > po2.Val {
                node := ListNode {
                    Val: po2.Val,
                }
                head.Next = &node
                head = head.Next
                po2 = po2.Next
            } else {
                node := ListNode {
                    Val: po1.Val,
                }
                head.Next = &node
                head = head.Next
                po1 = po1.Next
            }
        }
    }
    return ret.Next
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}