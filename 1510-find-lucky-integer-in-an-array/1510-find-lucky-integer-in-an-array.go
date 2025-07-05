func findLucky(arr []int) int {
    m := make(map[int]int)
    var res int = -1
    for _, v := range arr {
        if _, ok := m[v]; ok {
            m[v]++
        } else {
            m[v] = 1
        }
    }
    for i, f := range m {
        if i == f {
            res = max(res, f)
        }
    }
    return max(res, -1)
}