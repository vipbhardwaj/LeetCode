func makeFancyString(s string) string {
    if len(s) < 3 {
        return s
    }
    streak := 1
    var builder strings.Builder
    builder.WriteByte(s[0])
    for i:=1; i < len(s); i++ {
        if (s[i] == s[i-1]) {
            streak++
        } else {
            streak = 1
        }
        if streak < 3 {
            builder.WriteByte(s[i])
        }
    }
    return builder.String()
}