func largestAltitude(gain []int) int {
    var res, n int = max(0, gain[0]), len(gain)
    for i:=1; i < n; i++ {
        gain[i] += gain[i-1]
        res = max(res, gain[i])
    }
    return res
}