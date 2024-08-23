func strangePrinter(s string) int {
    var dp [100][100]int
    for j:=0; j < 100; j++ {
        for i:=0; i < 100; i++ {
            dp[j][i] = -1
        }
    }
    var dfs func(string, int, int) int
    dfs = func(s string, i int, j int) int {
        if i == j {
            return 1
        } else if dp[i][j] != -1 {
            return dp[i][j]
        }
        res := math.MaxInt
        for k:=i; k < j; k++ {
            res = min(res, dfs(s, i, k) + dfs(s, k+1, j))
        }
        if s[i] == s[j] {
            dp[i][j] = res - 1
            return res - 1
        }
        dp[i][j] = res
        return res
    }
    return dfs(s, 0, len(s)-1)
}