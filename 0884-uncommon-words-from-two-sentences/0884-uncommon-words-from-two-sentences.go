func uncommonFromSentences(s1 string, s2 string) []string {
    m := make(map[string]int)
    words1 := strings.Split(s1, " ")
    words2 := strings.Split(s2, " ")
    
    for _, word := range words1 {
        m[word]++
    }
    for _, word := range words2 {
        m[word]++
    }
    var res []string
    for word, count := range m {
        if count == 1 {
            res = append(res, word)
        }
    }
    return res
}