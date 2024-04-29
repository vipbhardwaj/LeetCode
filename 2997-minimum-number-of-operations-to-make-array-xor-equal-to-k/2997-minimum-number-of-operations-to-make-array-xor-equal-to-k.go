func minOperations(nums []int, k int) int {
    for _, num := range nums {
        k ^= num
    }
    
    ans := 0
    for k > 0 {
        ans += k%2
        k >>= 1
    }
    return ans
}