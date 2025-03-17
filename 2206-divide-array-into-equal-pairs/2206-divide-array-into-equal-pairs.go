func divideArray(nums []int) bool {
    m := make(map[int]bool)
    for _, n := range nums {
        if _, ok := m[n]; ok {
            m[n] = !m[n]
        } else {
            m[n] = false
        }
    }
    for _, b := range m {
        if !b {
            return false
        }
    }
    return true
}