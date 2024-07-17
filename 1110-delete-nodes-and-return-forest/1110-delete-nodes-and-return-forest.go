/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    m := make(map[int]int)
    for _, i := range to_delete {
        m[i] = i
    }
    r := []*TreeNode{}
    if _, ok := m[root.Val]; !ok {
        r = append(r, root)
    }
    r = dfs(root, m, r)
    // var res []*TreeNode
    // for _, i := range r {
    //     if _, ok := m[i.Val]; ok {
    //         continue
    //     }
    //     res = append(res, i)
    // }
    return r
}

func dfs(n *TreeNode, d map[int]int, r []*TreeNode) []*TreeNode {
    if n == nil || (n.Left==nil && n.Right==nil) {
        return r
    }
    
    if _, ok := d[n.Val]; ok {
        if n.Left != nil {
            if _, ok := d[n.Left.Val]; !ok {
                r = append(r, n.Left)
            }
        }
        if n.Right != nil {
            if _, ok := d[n.Right.Val]; !ok {
                r = append(r, n.Right)
            }
        }
    }
    
    r = dfs(n.Left, d, r)
    r = dfs(n.Right, d, r)
    
    if n.Left != nil {
        if _, ok := d[n.Left.Val]; ok {
            n.Left = nil
        }
    } 
    if n.Right != nil {
        if _, ok := d[n.Right.Val]; ok {
            n.Right = nil
        }
    }
    if _, ok := d[n.Val]; ok {
        if n.Left != nil {
            n.Left = nil
        }
        if n.Right != nil {
            n.Right = nil
        }
    }
    
    return r
} 