func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    slices.Sort(arr)
    ans, l := 1, len(arr)
    for i:=1; i<l; i++ {
        if arr[i] >= ans + 1 { ans += 1}
    }
    return ans
}