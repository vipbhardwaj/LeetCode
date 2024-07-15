/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    m := make(map[int]*TreeNode)
    notParent := make(map[*TreeNode]bool)
    var head *TreeNode
    for _, desc := range descriptions {
        p := desc[0]
        c := desc[1]
        l := desc[2]

        var pn *TreeNode
        if _, ok := m[p]; !ok {
            pn = &(TreeNode{p, nil, nil})
            m[p] = pn
        } else {
            pn = m[p]
        }
        if _, ok := notParent[pn]; ok && !notParent[pn] {
            notParent[pn] = false
        } else {
            notParent[pn] = true
        }

        var cn *TreeNode
        if _, ok := m[c]; !ok {
            cn = &(TreeNode{c, nil, nil})
            m[c] = cn
        } else {
            cn = m[c]
        }
        notParent[cn] = false

        if l==1 {
            pn.Left = cn
        } else {
            pn.Right = cn
        }
    }

    for k, v := range notParent {
        if v {
            head = k
            break
        }
    }
    return head
}