func findMin(nums []int) int {
    // O(n)
    var res int = math.MaxInt
    for _, n := range nums {
        res = min(res, n)
    }
    return res
}