func minimumCost(cost []int) int {
    sort.Slice(cost, func(i, j int) bool {
        return cost[i] > cost[j]
    })
    var res int
    for i, c := range cost {
        if (i+1) % 3 == 0 {
            continue
        }
        res += c
    }
    return res
}