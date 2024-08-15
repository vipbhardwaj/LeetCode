func lemonadeChange(bills []int) bool {
    var b5, b10 int
    for _, b := range bills {
        if b == 5 {
            b5++
            continue
        } else if b5 == 0 {
            return false
        }
        b5--
        if b == 20 {
            if b10 == 0 {
                if b5 < 2 {
                    return false
                }
                b5 -= 2
                continue
            }
            b10--
            continue
        }
        b10++
    }
    return true
}