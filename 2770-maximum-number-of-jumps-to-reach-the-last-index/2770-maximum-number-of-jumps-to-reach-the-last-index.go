func maximumJumps(nums []int, t int) int {
    var l int = len(nums)
    dp := make([]int, l)
    for i, _ := range dp {
        dp[i] = -1
    }
    dp[0] = 0

    for i:=0; i < l; i++ {
        for j:=0; j < i; j++ {
            if dp[j] != -1 && abs(nums[i], nums[j]) <= t {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
    }
    return dp[l-1]
}

func abs(i, j int) int {
    if i > j {
        return i - j
    }
    return j - i
}