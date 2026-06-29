func numOfStrings(patterns []string, word string) int {
    res := 0
    for _, s := range patterns {
        if strings.Contains(word, s) { res++ }
    }
    return res
}