func minFlips(s string) int {
    n := len(s)
    s = s + s

    alt1 := make([]byte, len(s))
    alt2 := make([]byte, len(s))

    for i := range s {
        if i%2 == 0 {
            alt1[i] = '0'
            alt2[i] = '1'
        } else {
            alt1[i] = '1'
            alt2[i] = '0'
        }
    }

    res := n
    diff1, diff2 := 0, 0
    l := 0

    for r := 0; r < len(s); r++ {
        if s[r] != alt1[r] { diff1++ }
        if s[r] != alt2[r] { diff2++ }

        if r-l+1 > n {
            if s[l] != alt1[l] { diff1-- }
            if s[l] != alt2[l] { diff2-- }
            l++
        }

        if r-l+1 == n {
            res = min(res, diff1, diff2)
        }
    }

    return res
}