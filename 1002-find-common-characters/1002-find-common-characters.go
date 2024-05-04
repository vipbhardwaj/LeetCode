func commonChars(words []string) []string {
    chars := make([]int, 26)
    for _, c := range words[0] {
        i := int(c - 'a')
        chars[i]++
    }
    for _, word := range words[1:] {
        char := make([]int, 26)
        for _, c := range word {
            i := int(c - 'a')
            char[i]++
        }
        for i:=0; i < 26; i++ {
            chars[i] = int(math.Min(float64(chars[i]), float64(char[i])))
        }
    }

    var ans []string
    for i, b := range chars {
        for b > 0 {
            ans = append(ans, string(i + 'a'))
            b--
        }
    }
    return ans
}