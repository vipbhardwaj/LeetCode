func isGood(nums []int) bool {
    m := make(map[int]int)
    limit := len(nums) - 1

    for _, n := range nums {
        if n > limit {
            return false
        }

        m[n]++
        val, ok := m[n]
        if ok && n == limit && val > 2 {
            return false
        }
        if ok && n < limit && val > 1 {
            return false
        }
    }
    return true
}