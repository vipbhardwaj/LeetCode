func helper(n int) bool {
    m := make([]int, 10)
    for n > 0 {
        m[n % 10]++
        n /= 10
    }
    for d:=0; d < 10; d++ {
        if m[d] > 0 && m[d] != d {
            return false
        }
    }
    return true
}

func nextBeautifulNumber(n int) int {
    for i := n+1; i <= 1224444; i++ {
        if helper(i) {
            return i
        }
    }
    return n
}