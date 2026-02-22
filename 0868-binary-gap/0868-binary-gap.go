func binaryGap(n int) int {
    var bits []bool
    for n > 0 {
        if n & 1 == 1 {
            bits = append(bits, true)
        } else {
            bits = append(bits, false)
        }
        n >>= 1
    }
    var l, res int = -1, 0
    for i, b := range bits {
        if b {
            if l == -1 {
                l = i
            } else {
                res = max(res, i - l)
                l = i
            }
        }
    }
    return res
}