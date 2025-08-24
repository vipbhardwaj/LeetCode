func longestSubarray(nums []int) int {
    var i, j, flag, res int = 0, 0, -1, -1
    for; j < len(nums); j++ {
        // for j < len(nums) && nums[j] == 1 {
        //     j++
        // }
        // flag = j
        for; j < len(nums) && nums[j] == 1; j++ {}
        if j == len(nums) {
            break
        }
        if flag == -1 {
            res = j-i-1
            flag = j
            continue
        }
        res = max(res, j-i-1)
        i = flag + 1
        flag = j
    }
    return max(res, j-i-1)
}