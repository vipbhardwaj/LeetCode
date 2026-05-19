func getCommon(nums1 []int, nums2 []int) int {
    m := make(map[int]bool)
    for _, n := range nums1 {
        m[n] = true
    }
    var res int = math.MaxInt
    for _, n := range nums2 {
        if _, ok := m[n]; ok {
            res = min(res, n)
        }
    }
    if res == math.MaxInt {
        return -1
    }
    return res
}