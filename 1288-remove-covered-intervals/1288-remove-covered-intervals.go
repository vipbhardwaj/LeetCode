func removeCoveredIntervals(intervals [][]int) int {
    m := make(map[[2]int]int)
    l := len(intervals)
    for _, i := range intervals {
        key := [2]int{i[0], i[1]}
        m[key]++
    }
    for i:=0; i<l; i++ {
        curr := intervals[i]
        currKey := [2]int{curr[0], curr[1]}
        for j:=0; j<l; j++ {
            if j == i { continue }
            target := intervals[j]
            if currKey[0] >= target[0] && currKey[1] <= target[1] {
                delete(m, currKey)
            }
        }
    }
    return len(m)
}