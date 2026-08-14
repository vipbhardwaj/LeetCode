func maximumLengthSubstring(s string) int {
    i, j, n, res := 0, 0, len(s), 0
    m := make(map[byte]int)
    for; i<n && j<n; j++ {
        m[s[j]]++
        for m[s[j]] > 2 {
            m[s[i]]--
            i++
        }
        res = max(res, j-i+1)
    }
    return res
}