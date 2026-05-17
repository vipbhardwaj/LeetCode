func canReach(arr []int, start int) bool {
    if arr[start] == 0 {
        return true
    }
    var q []int
    vis := make([]bool, len(arr))
    vis[start] = true
    q = enQ(arr, q, start)
    for pos:=0; pos < len(q); pos++ {
        curr := q[pos]
        if vis[curr] {
            continue
        }
        vis[curr] = true
        if arr[curr] == 0 {
            return true
        }
        q = enQ(arr, q, curr)
    }
    return false
}

func enQ(arr, q []int, i int) []int {
    if i - arr[i] >= 0 {
        q = append(q, i - arr[i])
    }
    if i + arr[i] < len(arr) {
        q = append(q, i + arr[i])
    }
    return q
}