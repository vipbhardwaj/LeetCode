func evaluate(s string, knowledge [][]string) string {
    m := make(map[string]string, len(knowledge))
    // m := make(map[string]string)
    for _, i := range knowledge { m[i[0]] = i[1] }

    n := len(s)
    var result, key strings.Builder
    inBrackets := false
    for i:=0; i<n; i++ {
        c := s[i]
        if c == '(' {
            inBrackets = true
        } else if c == ')' {
            k := key.String()
            if _, ok := m[k]; !ok { m[k] = "?" }
            result.WriteString(m[k])
            inBrackets = false
            key.Reset()
        } else {
            if !inBrackets {
                result.WriteByte(c)
            } else {
                key.WriteByte(c)
            }
        }
    }
    return result.String()
}