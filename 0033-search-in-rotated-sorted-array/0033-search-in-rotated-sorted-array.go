func search(nums []int, target int) int {
    var n int = len(nums)
    var lo, hi int = 0, n-1
    for lo < hi {
        var mid int = lo + (hi - lo) / 2
        if nums[mid] > nums[n-1] {
            lo = mid + 1
        } else {
            hi = mid
        }
    }

    var rot int = lo
    lo, hi = 0, n-1
    for lo <= hi {
        var mid int = lo + (hi - lo) / 2
        var midReal int = (mid + rot) % n
        if nums[midReal] == target {
            return midReal
        }
        if nums[midReal] < target {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return -1
}