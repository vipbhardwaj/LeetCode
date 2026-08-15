func longestSubsequence(nums []int) int {
    n, fullxor := len(nums), 0
    non_zero := true
    for _, i := range nums {
        fullxor ^= i
        if i > 0 {non_zero = false}
    }
    if fullxor > 0 {
        return n
    }
    if non_zero {
        return 0
    }
    return n-1
}