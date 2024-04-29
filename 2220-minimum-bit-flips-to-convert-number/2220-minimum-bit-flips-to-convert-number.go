func minBitFlips(start int, goal int) int {
    res := start ^ goal
    ans := 0
    for res > 0 {
        ans += res%2
        res >>= 1
    }
    return ans
}