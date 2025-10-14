func hasIncreasingSubarrays(nums []int, k int) bool {
    for i:=k; i <= len(nums)-k; i++ {
        var prev, next bool = true, true
        for j:=1; j < k; j++ {
            if nums[i-j] <= nums[i-j-1] {
                fmt.Println(nums[i-j], nums[i-j-1])
                prev = false
                break
            }
        }
        if !prev {
            continue
        }
        for j:=0; j < k-1; j++ {
            if nums[i+j] >= nums[i+j+1] {
                fmt.Println(nums[i+j], nums[i+j+1])
                next = false
                break
            }
        }
        if !next {
            continue
        }
        return true
    }
    return false
}