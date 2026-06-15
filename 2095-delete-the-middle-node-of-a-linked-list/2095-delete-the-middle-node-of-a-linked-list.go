/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return nil }

    temp := head
    prev := temp
    for i:=head; i != nil && i.Next != nil; i = i.Next.Next {
        prev = temp
        temp = temp.Next
    }
    prev.Next = temp.Next
    return head
}