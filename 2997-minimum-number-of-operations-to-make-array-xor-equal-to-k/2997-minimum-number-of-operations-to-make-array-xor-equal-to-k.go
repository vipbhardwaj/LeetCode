func minOperations(nums []int, k int) int {
    for _, num := range nums {
        k ^= num
    }
    
    ans := 0
    for k > 0 {
        if k & 1 == 1 {
            ans += 1
        }
        k >>= 1
    }
    return ans
}