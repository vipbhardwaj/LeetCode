func isAna(a, b string) bool {
    am, bm := make([]int, 26), make([]int, 26)
    for _, c := range a {
        am[c - 'a']++
    }
    for _, c := range b {
        bm[c - 'a']++
    }
    for i:=0; i < 26; i++ {
        if am[i] != bm[i] {
            return false
        }
    }
    return true
}

func removeAnagrams(words []string) []string {
    var res []string
    res = append(res, words[0])
    for i:=1; i < len(words); i++ {
        if isAna(res[len(res) - 1], words[i]) {
            continue
        }
        res = append(res, words[i])
    }
    return res
}