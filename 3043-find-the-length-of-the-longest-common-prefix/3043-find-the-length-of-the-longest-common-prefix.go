func longestCommonPrefix(arr1 []int, arr2 []int) int {
    prefixes := make(map[int]bool)
    for _, n := range arr1 {
        var x int = n
        for x > 0 {
            prefixes[x] = true
            x /= 10
        }
    }

    var ans int
    for _, n := range arr2 {
        var x int = n
        var len int = digits(n)
        for x > 0 {
            if _, ok := prefixes[x]; ok {
                ans = max(ans, len)
                break
            }
            x /= 10
            len--
        }
    }
    return ans
}

func digits(x int) int {
    var cnt int
    for x > 0 {
        cnt++
        x /= 10
    }
    return cnt
}