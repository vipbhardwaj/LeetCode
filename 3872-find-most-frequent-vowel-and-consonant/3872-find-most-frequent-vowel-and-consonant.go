func maxFreqSum(s string) int {
    m := make(map[rune]int)
    for _, c := range s {
        m[c]++
    }
    var v, c int
    for k, i := range m {
        if k == 'a' || k == 'e' || k == 'i' || k == 'o' || k == 'u' {
            v = max(v, i)
        } else {
            c = max(c, i)
        }
    }
    return v + c
}