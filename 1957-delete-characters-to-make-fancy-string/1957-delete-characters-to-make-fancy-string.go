func makeFancyString(s string) string {
    if len(s) < 3 {
        return s
    }

    var builder strings.Builder
    builder.WriteByte(s[0])
    builder.WriteByte(s[1])
    prev, curr := false, false
    if s[0] == s[1] {
        prev = true
    }
    for i:=2; i < len(s); i++ {
        if s[i] == s[i-1] {
            curr = true
        } else {
            curr = false
        }
        if curr && prev {
            prev = curr
            continue
        }
        prev = curr
        builder.WriteByte(s[i])
    }
    return builder.String()
}