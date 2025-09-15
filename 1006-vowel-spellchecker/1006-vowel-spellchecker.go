func spellchecker(wordlist []string, queries []string) []string {
    m := make(map[string]bool)
    c := make(map[string]string)
    v := make(map[string]string)
    for _, s := range wordlist {
        m[s] = true
        var ss string = strings.ToLower(s)
        if _, ok := c[ss]; !ok {
            c[ss] = s
        }
        var sss string = mask(ss)
        if _, ok := v[sss]; !ok {
            v[sss] = s
        }
    }

    var res []string
    for _, s := range queries {
        if _, ok := m[s]; ok {
            res = append(res, s)
            continue
        }
        var ss string = strings.ToLower(s)
        if k, ok := c[ss]; ok {
            res = append(res, k)
            continue
        }
        if k, ok := v[mask(ss)]; ok {
            res = append(res, k)
            continue
        }
        res = append(res, "")
    }
    return res
}

func mask(s string) string {
    var res strings.Builder
    for _, c := range s {
        if isVowel(c) {
            res.WriteRune('a')
            continue
        }
        res.WriteRune(c)
    }
    return res.String()
}

func isVowel(c rune) bool {
    switch c {
    case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
        return true
    }
    return false
}