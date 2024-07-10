func minOperations(logs []string) int {
    var depth int
    for _, l := range logs {
        if l == "../" {
            if depth <= 0 {
                continue
            }
            depth -= 1
        } else if l == "./" {
            continue
        } else {
            depth += 1
        }
    }
    return depth
}