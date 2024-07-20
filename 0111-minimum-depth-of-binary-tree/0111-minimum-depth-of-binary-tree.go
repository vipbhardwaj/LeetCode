/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return bfs(root, 1)
    return dfs(root, 1)
}

func bfs(r *TreeNode, d int) int {
    q := []*TreeNode{r}
    f:=0
    for f < len(q) {
        s := len(q)
        for; f<s; f++ {
            n := q[f]
            if n.Left == nil && n.Right == nil {
                return d
            }
            if n.Left != nil {
                q = append(q, n.Left)
            }
            if n.Right != nil {
                q = append(q, n.Right)
            }
        }
        d++
    }
    return d
}

func dfs(n *TreeNode, d int) int {
    if n.Left == nil && n.Right == nil {
        return d
    }
    
    ld, rd := math.MaxInt, math.MaxInt
    if n.Left != nil {
        ld = dfs(n.Left, d+1)
    }
    if n.Right != nil {
        rd = dfs(n.Right, d+1)
    }
    
    return min(ld, rd)
}