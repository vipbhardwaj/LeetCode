func numberOfSpecialChars(word string) int {
    m := make(map[rune]int)
    M := make(map[rune]int)
    var res int
    for i, c := range word {
        if c >= 97 {
            m[c] = i
        } else if _, ok := M[c]; !ok {
            M[c] = i
        }
    }
    for c, v := range m {
        if i, ok := M[c-32]; ok && v < i {
            res++
        }
    }
    return res
}