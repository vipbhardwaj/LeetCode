func minimumCost(n int, edges [][]int, query [][]int) []int {
    var adj = make([][][]int, n)
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], []int{e[1], e[2]})
        adj[e[1]] = append(adj[e[1]], []int{e[0], e[2]})
    }

    var comp = make([]int, n)
    var vis = make([]bool, n)
    var compCost, res []int
    var compId int
    for i := range n {
        if !vis[i] {
            compCost = append(compCost, getCost(i, adj, &vis, comp, compId))
            compId++
        }
    }

    for _, q := range query {
        if comp[q[0]] != comp[q[1]] {
            res = append(res, -1)
            continue
        }
        res = append(res, compCost[comp[q[0]]])
    }
    return res
}

func getCost(n int, adj [][][]int, v *[]bool, comp []int, compId int) int {
    comp[n] = compId
    (*v)[n] = true
    cost := math.MaxInt
    for _, m := range adj[n] {
        cost &= m[1]
        if !(*v)[m[0]] {
            cost &= getCost(m[0], adj, v, comp, compId)
        }
    }
    return cost
}