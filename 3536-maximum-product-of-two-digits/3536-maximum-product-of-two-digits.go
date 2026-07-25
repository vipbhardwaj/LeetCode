func maxProduct(n int) int {
    prev, curr := n%10, 0
    res := 0
    for n > 0 {
        n /= 10
        curr = n % 10
        res = max(res, prev * curr)
        prev = max(prev, curr)
    }
    return res
}