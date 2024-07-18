/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countPairs(root *TreeNode, distance int) int {
    c := 0
    dfs(root, distance, &c)
    return c
}

func dfs(n *TreeNode, d int, c *int) []int {
    if n == nil {
        return []int{}
    } else if n.Left == nil && n.Right == nil {
        return []int{1}
    }
    
    l := dfs(n.Left, d, c)
    r := dfs(n.Right, d, c)
    x := make([]int, len(l) + len(r))
    
    a, b := len(l)-1, 0
    for; a >= 0; a-- {
        for; b < len(r) && l[a] + r[b] <= d; b++ {}
        *c += b
    }
    
    i, j, k := 0, 0, 0
    for; i < len(l) && j < len(r); k++ {
        if l[i] < r[j] {
            x[k] = l[i]+1
            i++
        } else {
            x[k] = r[j]+1
            j++
        }
        // fmt.Println("hi")
    }
    // fmt.Println("x:", x, ", k:" , k)
    for; i < len(l); k++ {
        // fmt.Print(k,",")
        x[k] = l[i]+1
        i++
    }
    for; j < len(r); k++ {
        x[k] = r[j]+1
        j++
    }
    return x
}