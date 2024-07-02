func minOperations(nums []int, k int) int {
    // Initialize a list to track the elements found
    var found []bool
    for range(k) {
        found = append(found, false)
    }
    
    // Start tracking the elements
    var ops = 0
    var marked = 0
    for i := len(nums)-1; i >= 0; i-- {
        ops++
        if nums[i]-1 >= k {
            continue
        }
        if !found[nums[i]-1] {
            found[nums[i]-1] = true
            marked++
        }
        if marked == k {
            break
        }
    }
    return ops
}