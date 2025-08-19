func zeroFilledSubarray(nums []int) int64 {
    var res, currZeroes int64
    for _, v := range nums {
        if  v != 0 {
            currZeroes = 0
            continue
        }
        res += currZeroes + 1
        currZeroes++
    }
    return res
}