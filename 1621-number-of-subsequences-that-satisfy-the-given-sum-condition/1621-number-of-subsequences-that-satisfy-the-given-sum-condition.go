func numSubseq(nums []int, target int) int {
    sort.Ints(nums)
    pow2 := make([]int, len(nums))
    pow2[0] = 1
    for i:=1; i < len(nums); i++ {
        pow2[i] = int(math.Mod(float64(pow2[i-1] * 2), float64(1000000007)))
    }
    var i, j, ans int
    j = len(nums) - 1
    for i <= j {
        if nums[i] + nums[j] <= target {
            ans = int(math.Mod(float64(ans + pow2[j-i]), float64(1000000007)))
            i++
        } else {
            j--
        }
    }
    return ans
}

func MOD(x, y int) int {
    if y == -1 {
        return int(math.Mod(float64(x), float64(1000000007)))
    }
    return int(math.Mod(float64(x) + math.Mod(float64(y), float64(1000000007)), float64(1000000007)))
}