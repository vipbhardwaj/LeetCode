func rotateString(s string, goal string) bool {
    if s == goal {
        return true
    }
    var l int = len(s)
    var ss strings.Builder
    ss.WriteString(s)
    ss.WriteString(s)
    s = ss.String()

    for i := 1; i < l; i++ {
        if makeStr(i, s, l) == goal {
            return true
        }
    }
    return false
}

func makeStr(i int, s string, n int) string {
    var ss strings.Builder
    for j := 0; j < n; j++ {
        ss.WriteByte(s[i+j])
    }
    return ss.String()
}