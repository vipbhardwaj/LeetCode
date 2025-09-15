func canBeTypedWords(text string, brokenLetters string) int {
    m := make(map[rune]bool)
    for _, c := range brokenLetters {
        m[c] = true
    }

    var words []string = strings.Split(text, " ")
    var res int
    for _, w := range words {
        var flag bool = true
        for _, c := range w {
            if _, ok := m[c]; ok {
                flag = false
                break
            }
        }
        if flag {
            res++
        }
    }
    return res
}