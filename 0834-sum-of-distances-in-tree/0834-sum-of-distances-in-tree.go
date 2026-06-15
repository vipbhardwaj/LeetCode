func sumOfDistancesInTree(n int, edges [][]int) []int {
    m := make(map[int][]int)
    count, res := make([]int, n), make([]int, n)
    for i:=0; i < n; i++ { count[i] = 1 }

    for _, e := range edges {
        m[e[0]] = append(m[e[0]], e[1])
        m[e[1]] = append(m[e[1]], e[0])
    }
    postOrder(m, 0, -1, count, res)
    preOrder(m, 0, -1, count, res)
    return res
}

func postOrder(m map[int][]int, node, parent int, count, res []int) {
    for _, child := range m[node] {
        if child == parent { continue }
        postOrder(m, child, node, count, res)
        count[node] += count[child]
        res[node] += res[child] + count[child]
    }
}

func preOrder(m map[int][]int, node, parent int, count, res []int) {
    for _, child := range m[node] {
        if child == parent { continue }
        res[child] = res[node] - count[child] + len(m) - count[child]
        preOrder(m, child, node, count, res)
    }
}