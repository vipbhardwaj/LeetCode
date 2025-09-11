func sortVowels(s string) string {
    var allVowels []rune
    for _, c := range s {
        if !isVowel(c) {
            continue
        }
        allVowels = append(allVowels, c)
    }
    sort.Slice(allVowels, func(i, j int) bool {
        return allVowels[i] < allVowels[j]
    })

    var res strings.Builder
    var i int
    for _, c := range s {
        if isVowel(c) {
            res.WriteRune(allVowels[i])
            i++
            continue
        }
        res.WriteRune(c)
    }
    return res.String()

}

func isVowel(c rune) bool {
    if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
        return true
    }
    if c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
        return true
    }
    return false
}