func rotateString(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }

    s = s + s
    if strings.Contains(s, goal) {
        return true
    }
    return false
}