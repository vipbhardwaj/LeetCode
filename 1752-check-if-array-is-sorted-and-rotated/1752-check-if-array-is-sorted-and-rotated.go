func check(nums []int) bool {
    var n int = len(nums)
    var drop bool
    for i:=0; i < n; i++ {
        if nums[i] > nums[(i+1)%n] {
            if drop {
                return false
            }
            drop = true
        }
    }
    return true
}