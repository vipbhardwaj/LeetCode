func isOneBitCharacter(bits []int) bool {
    var i int
    for i < len(bits) - 1 {
        if bits[i] == 1 {
            i += 2
        } else {
            i++
        }
    }
    return i == len(bits) - 1
}