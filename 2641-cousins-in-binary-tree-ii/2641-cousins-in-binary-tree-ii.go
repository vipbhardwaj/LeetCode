/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
    cousin := make(map[int]int)
    sibling := make(map[*TreeNode]int)
    dfs(root, cousin, 0)
    dfs2(root, sibling, 0)
    dfs3(root, sibling, cousin, root, 0)
    // for i, j := range cousin {
    //     fmt.Println("Depth:", i, ", Cousin Sum:", j)
    // }
    // fmt.Println("---------------------")
    // for i, j := range sibling {
    //     fmt.Println("Depth:", i, ", Sibling Sum:", j)
    // }
    return root
}

func dfs3(n *TreeNode, s map[*TreeNode]int, c map[int]int, p *TreeNode, d int) {
    if n == nil {
        return
    }
    if c[d] == 0 {
        n.Val = 0
    } else {
        n.Val = c[d] - s[p]
    }
    dfs3(n.Left, s, c, n, d+1)
    dfs3(n.Right, s, c, n, d+1)
}

func dfs2(n *TreeNode, m map[*TreeNode]int, d int) {
    if n == nil {
        return
    }
    if n.Left != nil {
        m[n] = n.Left.Val
    }
    if n.Right != nil {
        m[n] += n.Right.Val
    }
    dfs2(n.Left, m, d+1)
    dfs2(n.Right, m, d+1)
}

func dfs(n *TreeNode, m map[int]int, d int) {
    if n == nil {
        return
    }
    if d < 2 {
        m[d] = 0
    } else {
        if _, ok := m[d]; ok {
            m[d] += n.Val
        } else {
            m[d] = n.Val
        }
    }
    dfs(n.Left, m, d+1)
    dfs(n.Right, m, d+1)
}