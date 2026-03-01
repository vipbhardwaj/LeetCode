func minPartitions(n string) int {
    var res int
    for _, c := range n {
        res = max(res, int(c-'0'))
    }
    return res
}