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
    return dfs(root, 1)
}

func dfs(n *TreeNode, d int) int {
    if n.Left == nil && n.Right == nil {
        return d
    }
    
    ld, rd := 100000000, 100000000
    if n.Left != nil {
        ld = dfs(n.Left, d+1)
    }
    if n.Right != nil {
        rd = dfs(n.Right, d+1)
    }
    
    return min(ld, rd)
}