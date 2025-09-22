func maxFrequencyElements(nums []int) int {
    m := make(map[int]int)
    var maxFreq, res int
    for _, e := range nums {
        m[e]++
        maxFreq = max(maxFreq, m[e])
    }
    for _, f := range m {
        if f == maxFreq {
            res += maxFreq
        }
    }
    return res
}