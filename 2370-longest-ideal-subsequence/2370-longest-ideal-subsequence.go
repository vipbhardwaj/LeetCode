func longestIdealString(s string, k int) int {
    dp := make([]int, 26)

    for _, c := range s {
        c := int(c - 'a')
        longest := 1
        for i, x := range dp {
            if absolut(c, i) <= k {
                longest = max(longest, 1 + x)
            }
        }
        dp[c] = max(dp[c], longest)
    }

    res := 0
    for _, x := range dp {
        res = max(x, res)
    }
    return res
}

func absolut(c int, i int) int {
    if c > i {
        return c-i
    } else {
        return i-c
    }
}