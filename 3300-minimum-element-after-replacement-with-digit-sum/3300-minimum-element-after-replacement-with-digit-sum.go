func minElement(nums []int) int {
    var res int = 100000000
    for _, n := range nums {
        res = min(res, digitSum(n)) 
    }
    return res
}

func digitSum(n int) int {
    var res int
    for n > 0 {
        res += n%10
        n /= 10
    }
    return res
}