func maximumCandies(candies []int, k int64) int {
    var most int
    for _, c := range candies {most = max(most, c)}
    l, r, m := 0, most, 0
    for l < r {
        m = (l+r+1) / 2
        if pilesOf(candies, m) >= k {
            l = m
        } else {
            r = m - 1
        }
    }
    return l
}

func pilesOf(candies []int, m int) int64 {
    var g int
    for _, c := range candies {g += c / m}
    return int64(g)
}