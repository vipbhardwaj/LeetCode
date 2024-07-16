/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getDirections(root *TreeNode, startValue int, destValue int) string {
    ups, path1, _ := dfs(root, startValue, []rune{}, []int{})
    downs, path2, _ := dfs(root, destValue, []rune{}, []int{})
    
    for i, j := 0, len(ups)-1; i < j; i, j = i+1, j-1 {
        ups[i], ups[j] = ups[j], ups[i]
    }
    for i, j := 0, len(downs)-1; i < j; i, j = i+1, j-1 {
        downs[i], downs[j] = downs[j], downs[i]
    }
    // s1 := string(ups)

    // fmt.Println("UPS:", string(ups))
    // fmt.Println("path1:", path1)
    // fmt.Println("DOWNS:", string(downs))
    // fmt.Println("path2:", path2)
    
    i := 1
    for ; i < len(path1) && i < len(path2); i++ {
        if path1[i] != path2[i] {
            break
        }
    }
    // fmt.Println(i)
    ups = ups[i-1:]
    downs = downs[i-1:]
    // fmt.Println(strings.Repeat("U", len(ups)))
    // fmt.Println(string(downs))
    return strings.Repeat("U", len(ups)) + string(downs)
}

func dfs(n *TreeNode, v int, arr []rune, path []int) ([]rune, []int, bool) {
    if n == nil {
        return arr, path, false
    }
    
    path = append(path, n.Val)
    
    
    if n.Val == v {
        return arr, path, true
    }
    
    larr, lpath, found := dfs(n.Left, v, arr, path)
    if found {
        return append(larr, 'L'), lpath, true
    }
    rarr, rpath, found := dfs(n.Right, v, arr, path)
    if found {
        return append(rarr, 'R'), rpath, true
    }
    
    return arr, path, false
}