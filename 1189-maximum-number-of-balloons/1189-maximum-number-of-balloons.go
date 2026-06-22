func maxNumberOfBalloons(text string) int {
    m := make(map[rune]int)
    need := []rune{'b', 'a', 'l', 'o', 'n'}
    for _, c := range text {
        if slices.Contains(need, c) { m[c]++ }
    }
    fmt.Println(m)
    m['l'] /= 2
    m['o'] /= 2
    return min(m['a'], min(m['b'], min(m['n'], min(m['l'], m['o']))))
}