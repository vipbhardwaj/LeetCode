func doesAliceWin(s string) bool {
    var vCount int
    for _, c := range s {
        if isVowel(c) {
            vCount++
        }
    }
    return vCount != 0
}

func isVowel(c rune) bool {
    if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
        return true
    }
    return false
}