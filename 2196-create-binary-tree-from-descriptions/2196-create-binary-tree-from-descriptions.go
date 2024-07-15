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
        // fmt.Println(p)
        // fmt.Println(c)
        
        var pn *TreeNode
        if _, ok := m[p]; !ok {
            pn = &(TreeNode{p, nil, nil})
            m[p] = pn
        } else {
            pn = m[p]
        }
        // fmt.Println(pn.Val)
        if _, ok := notParent[pn]; ok && !notParent[pn] {
            notParent[pn] = false
        } else {
            notParent[pn] = true
        }
        // fmt.Println("Found:", p, "with state:", notParent[pn])
        var cn *TreeNode
        if _, ok := m[c]; !ok {
            cn = &(TreeNode{c, nil, nil})
            m[c] = cn
        } else {
            cn = m[c]
        }
        // fmt.Println(cn.Val)
        notParent[cn] = false
        // fmt.Println(c, "is not parent")
        
        if l==1 {
            pn.Left = cn
        } else {
            pn.Right = cn
        }
    }
    for k, v := range notParent {
        // fmt.Println(v)
        if v {
            head = k
            // fmt.Println(head.Val)
            break
        }
    }
    return head
}