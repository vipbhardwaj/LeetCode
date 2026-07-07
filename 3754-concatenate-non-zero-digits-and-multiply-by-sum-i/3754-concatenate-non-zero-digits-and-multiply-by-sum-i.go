func sumAndMultiply(n int) int64 {
    x, sum := 0, 0
    for n > 0 {
        d := n % 10
        n /= 10
        if d == 0 { continue }
        x *= 10
        x += d
        sum += d
    }
    xx := 0
    for x > 0 {
        xx *= 10
        xx += (x % 10)
        x /= 10
    }
    return int64(xx) * int64(sum)
}