func countCompleteComponents(n int, edges [][]int) int {
    var adj = make([][]int, n)
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], e[1])
        adj[e[1]] = append(adj[e[1]], e[0])
    }
    v := make(map[int]bool)
    var res int
    for i := range n {
        if v[i] {continue}
        var compInfo = []int{0, 0}
        dfs(i, adj, v, &compInfo)
        if compInfo[0] * (compInfo[0]-1) == compInfo[1] {res++}
    }
    return res
}

func dfs(n int, adj [][]int, v map[int]bool, c *[]int) {
    v[n] = true
    (*c)[0]++
    (*c)[1] += len(adj[n])
    for _, j := range adj[n] {
        if v[j] {continue}
        dfs(j, adj, v, c)
    }
}