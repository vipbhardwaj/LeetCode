/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    prev := head
    curr := head.Next
    
    for curr != nil {
        sum := 0
        for curr.Val != 0 {
            sum += curr.Val
            curr = curr.Next
        }
        prev.Val = sum
        prev.Next = curr
        curr = curr.Next
        if curr != nil {
            prev = prev.Next
        } else {
            prev.Next = nil
        }
    }
    return head
}