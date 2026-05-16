func findMin(nums []int) int {
    var res int = math.MaxInt
    for _, n := range nums {
        res = min(res, n)
    }
    return res
}