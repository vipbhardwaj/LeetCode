func maxIceCream(costs []int, coins int) int {
    costs = countingSort(costs)
    res := 0
    for _, v := range costs {
        if coins < v { break }
        coins -= v
        res++
    }
    return res
}

func countingSort(arr []int) []int {
    n, m := len(arr), 0
    for _, v := range arr { m = max(v, m) }
    freq, res := make([]int, m+1), make([]int, n)
    for _, v := range arr { freq[v]++ }

    // PrefSum indices arr
    for i:=1; i<=m; i++ { freq[i] += freq[i-1] }
    for _, v := range arr {
        res[freq[v]-1] = v
        freq[v]--
    }
    return res
}