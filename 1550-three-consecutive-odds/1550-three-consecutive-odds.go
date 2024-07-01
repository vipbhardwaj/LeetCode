func threeConsecutiveOdds(arr []int) bool {
    for i := range(len(arr)-2) {
        if arr[i] % 2 == 0 {
            continue
        }
        if arr[i+1] % 2 == 0 {
            continue
        }
        if arr[i+2] % 2 == 0 {
            continue
        }
        return true
    }
    return false
}