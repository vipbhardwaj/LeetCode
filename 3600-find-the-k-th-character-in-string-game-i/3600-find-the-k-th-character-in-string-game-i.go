func kthCharacter(k int) byte {
    var s strings.Builder
    s.WriteRune('a')

    // var res = 'a'
    var ss = "a"
    for len(ss) < k {
        for i:=0; i < len(ss); i++ {
            s.WriteByte(ss[i] + 1)
        }
        ss = s.String()
    }
    return ss[k-1]
}