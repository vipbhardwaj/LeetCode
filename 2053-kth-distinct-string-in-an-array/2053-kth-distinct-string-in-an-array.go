func kthDistinct(arr []string, k int) string {
    seen := make(map[string]bool)
    for _, s := range arr {
        if _, ok := seen[s]; ok {
            seen[s] = true
        } else {
            seen[s] = false
        }
    }

    for _, s := range arr {
        if seen[s] {
            continue
        }
        k--
        if k == 0 {
            return s
        }
    }
    return ""
}