func zeroFilledSubarray(nums []int) int64 {
    nums = append(nums, 1)
    var sum, res int64 = 0, 0

    for _, n := range nums {
        if n == 0 {
            sum++
        } else {
            res += sum * (sum+1) / 2
            sum = 0
        }
    }
    return res
}