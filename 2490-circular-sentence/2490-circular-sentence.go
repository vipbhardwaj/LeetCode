func isCircularSentence(s string) bool {
    var last byte
    for i:=0; i < len(s); i++ {
        if s[i] != ' ' {
            last = s[i]
            continue
        }
        if i < len(s)-1 && s[i-1] != s[i+1] {
            return false
        }
    }
    if s[0] != last {
        return false
    }
    return true
}