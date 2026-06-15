/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil { return nil }

    temp := head
    i := head.Next.Next
    for i != nil && i.Next != nil {
        temp = temp.Next
        i = i.Next.Next
    }
    temp.Next = temp.Next.Next
    return head
}