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
        var nodes_edges = []int{0, 0}
        dfs(i, adj, v, nodes_edges)
        if nodes_edges[0] * (nodes_edges[0]-1) == nodes_edges[1] {res++}
    }
    return res
}

func dfs(node int, adj [][]int, v map[int]bool, nodes_edges []int) {
    v[node] = true
    nodes_edges[0]++
    nodes_edges[1] += len(adj[node])
    for _, j := range adj[node] {
        if v[j] {continue}
        dfs(j, adj, v, nodes_edges)
    }
}