func zeroFilledSubarray(nums []int) int64 {
    sum := 0
    var res int64
    for _, n := range nums {
        if n == 0 {
            sum += 1
        } else {
            res += int64(sum * (sum+1) / 2)
            sum = 0
        }
    }
    return res + int64(sum * (sum+1) / 2)
}