func maxUniqueSplit(s string) int {
    seen := make(map[string]bool)
    return backtrack(0, s, seen)
}

func backtrack(start int, s string, seen map[string]bool) int {
    if start == len(s) {
        return 0
    }
    var maxSplits int
    for end := start+1; end <= len(s); end++ {
        subs := s[start : end]
        if _, ok := seen[subs]; !ok {
            seen[subs] = true
            maxSplits = max(maxSplits, 1 + backtrack(end, s, seen))
            delete(seen, subs)
        }
    }
    return maxSplits
}