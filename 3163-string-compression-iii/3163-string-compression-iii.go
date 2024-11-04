func compressedString(word string) string {
    var builder strings.Builder
    count := 1
    for i:=0; i < len(word)-1; i++ {
        if word[i] == word[i+1] && count < 9 {
            count++
        } else {
            builder.WriteString(strconv.Itoa(count))
            builder.WriteByte(word[i])
            count = 1
        }
    }
    if count > 0 {
        builder.WriteString(strconv.Itoa(count))
        builder.WriteByte(word[len(word)-1])
    }
    return builder.String()
}