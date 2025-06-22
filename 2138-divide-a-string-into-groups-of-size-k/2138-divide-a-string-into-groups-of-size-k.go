func divideString(s string, k int, fill byte) []string {
    var res []string
    var ss strings.Builder
    var temp = k
    for _, c := range s {
        temp--
        ss.WriteRune(c)
        if temp == 0 {
            res = append(res, ss.String())
            temp = k
            ss.Reset()
        }
    }
    if ss.Len() == 0 {
        return res
    }
    temp = k - ss.Len()
    for temp > 0 {
        ss.WriteByte(fill)
        temp--
    }
    res = append(res, ss.String())
    return res
}