func buildArray(target []int, n int) []string {
    var ops []string
    j := 0
    for i := 1; i <= n; i++ {
        ops = append(ops, "Push")
        if i == target[len(target)-1] {
            break
        }
        if i != target[j] {
            ops = append(ops, "Pop")
        } else {
            j++
        }
    }
    return ops
}