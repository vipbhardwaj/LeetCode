func maxRotateFunction(nums []int) int {
    var res, t, s int
    for i, n := range nums {
        t += i * n
        s += n
    }
    res = t
    var l int = len(nums)
    for i:=1; i < l; i++ {
        t += (nums[i-1] * (l-1)) - (s - nums[i-1])
        res = max(res, t)
    }
    return res
}