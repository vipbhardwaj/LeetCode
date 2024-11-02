func isCircularSentence(s string) bool {
    // var last byte
    for i:=0; i < len(s); i++ {
        if s[i] != ' ' {
            continue
        }
        if s[i-1] != s[i+1] {
            return false
        }
    }
    if s[0] != s[len(s)-1] {
        return false
    }
    return true
}