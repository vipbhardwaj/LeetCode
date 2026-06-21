func processStr(s string, k int64) byte {
    n := int64(0)
    for _, c := range s {
        switch c {
        case '#':
            n *= 2
        case '*':
            if n > 0 { n-- }
        case '%':
        default:
            n++
        }
    }
    if k >= n { return '.' }

    l := len(s)
    for i:=l-1; i>=0; i-- {
        switch s[i] {
        case '%':
            k = n-k-1
        case '#':
            if k >= (n+1)/2 { k -= n/2 }
            n = (n+1)/2
        case '*':
            n++
        default:
            if k+1 == n {
				return s[i]
			}
			n--
        }
    }
    return '.'
}