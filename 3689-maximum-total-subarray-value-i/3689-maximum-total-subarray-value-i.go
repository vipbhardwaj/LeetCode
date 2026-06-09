func maxTotalValue(nums []int, k int) int64 {
    if len(nums) == 1 {
        return 0
    }
    var alpha, beta int = 0, 100000000
    for _, n := range nums {
        alpha = max(alpha, n)
        beta = min(beta, n)
    }
    return int64(k) * int64(alpha - beta)
}