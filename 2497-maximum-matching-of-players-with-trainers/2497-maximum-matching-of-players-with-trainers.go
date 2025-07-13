func matchPlayersAndTrainers(p []int, t []int) int {
    sort.Ints(p)
    sort.Ints(t)

    var res, ti, pi, i, j int = 0, len(t), len(p), 0, 0
    for i < pi && j < ti {
        if p[i] <= t[j] {
            res++
            i++
        }
        j++
    }
    return res
}