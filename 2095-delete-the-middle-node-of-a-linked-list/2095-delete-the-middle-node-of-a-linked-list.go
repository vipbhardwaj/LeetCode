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
    for i:=head; i != nil; i = i.Next.Next {
        if i.Next == nil { break }
        prev = temp
        temp = temp.Next
    }
    prev.Next = temp.Next
    return head
}