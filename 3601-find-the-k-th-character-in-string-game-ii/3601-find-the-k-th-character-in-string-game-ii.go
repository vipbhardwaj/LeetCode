func kthCharacter(k int64, ops []int) byte {
    var s int64 = 1
    var op int
    for s < k {
        s *= 2
        op++
    }
    
    var shifts int
    for s > 1 {
        op--
        s /= 2
        if k > s {
            k -= s
            shifts += ops[op]
        }
    }
    return byte((('a' - 'a' + shifts) % 26) + 'a')
}