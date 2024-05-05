func zeroFilledSubarray(nums []int) int64 {
    var sum int64
    var res int64
    for _, n := range nums {
        if n == 0 {
            sum += 1
        } else {
            res += sum * (sum+1) / 2
            sum = 0
        }
    }
    return res + (sum * (sum+1) / 2)
}