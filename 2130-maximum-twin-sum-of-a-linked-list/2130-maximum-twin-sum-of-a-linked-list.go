/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    var twinArr []int
    for head != nil {
        twinArr = append(twinArr, head.Val)
        head = head.Next
    }

    var res int
    var n int = len(twinArr)
    var l int = n/2 -1
    for i:=0; i <= l; i++ {
        res = max(res, twinArr[i] + twinArr[n-1-i])
    }
    return res
}