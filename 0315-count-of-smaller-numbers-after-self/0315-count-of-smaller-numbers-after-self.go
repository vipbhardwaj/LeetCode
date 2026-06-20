func countSmaller(nums []int) []int {
    n := len(nums)
    pairs := make([][]int, n)
    temp := make([][]int, n)
    ans := make([]int, n)
    for i, v := range nums {
        pairs[i] = []int{v, i}
        temp[i] = []int{v, i}
    }
    merge_sort(0, n-1, ans, pairs, temp)
    return ans
}

func merge_sort(s, e int, ans []int, pairs, temp [][]int) {
    if s >= e { return }
    m := s + (e - s) / 2
    merge_sort(s, m, ans, pairs, temp)
    merge_sort(m+1, e, ans, pairs, temp)

    l, r, i, nRightLessThanLeft := s, m + 1, s, 0
    for l <= m && r <= e {
        if pairs[l][0] <= pairs[r][0] {
            ans[pairs[l][1]] += nRightLessThanLeft
            temp[i] = pairs[l]
            l++
        } else {
            temp[i] = pairs[r]
            nRightLessThanLeft++
            r++
        }
        i++
    }
    for l <= m {
        ans[pairs[l][1]] += nRightLessThanLeft
        temp[i] = pairs[l]
        i++
        l++
    }
    for r <= e {
        temp[i] = pairs[r]
        i++
        r++
    }
    for j:=s; j<=e; j++ { pairs[j] = temp[j] }
}