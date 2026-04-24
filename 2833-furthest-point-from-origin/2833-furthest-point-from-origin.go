func furthestDistanceFromOrigin(moves string) int {
    var r, l, any int
    for _, c := range moves {
        if c == 'R' {
            r++
        } else if c == 'L' {
            l++
        } else {
            any++
        }
    }
    return max(l, r) - min(l, r) + any
}