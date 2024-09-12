func countConsistentStrings(allowed string, words []string) int {
    chars := make([]bool, 26)
    for _, c := range allowed {
        chars[c-'a'] = true
    }

    var res int // 0
    for _, w := range words {
        for i, c := range w {
            if !chars[c-'a'] {
                break 
            }
            if i == len(w)-1 {
                res++
            }
        }
    }
    return res
}