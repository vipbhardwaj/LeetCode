func minFlips(s string) int {
    res := I(s)
    if s[0] == s[len(s)-1] {
        return res
    }
    var i int = 1
    var b strings.Builder
    b.Grow(len(s))
    for; i < len(s) - 1; i++ {
        if s[i] == s[i-1] {
            break
        }
    }
    for ; i < len(s); i++ {
        b.WriteByte(s[i % len(s)])
    }
    return min(res, I(b.String()))
}

func I(s string) int {
    var res int
    prev := s[0]
    run := 1
    for i:=1; i < len(s); i++ {
        curr := s[i]
        if curr != prev {
            res += run / 2
            run = 1
        } else {
            run++
        }
        prev = curr
        if i >= len(s) / 2 {
            
        }
    }
    fmt.Println(res)
    return res + run / 2
}