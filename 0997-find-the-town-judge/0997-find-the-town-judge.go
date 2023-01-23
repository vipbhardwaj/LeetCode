func findJudge(n int, trust [][]int) int {
    t := make([]int, n)
    
    for _, i := range trust {
        t[i[0] - 1]--
        t[i[1] - 1]++
    }
    for i, v := range t {
        if v == n - 1 {
            return i + 1
        }
    }
    return -1
}