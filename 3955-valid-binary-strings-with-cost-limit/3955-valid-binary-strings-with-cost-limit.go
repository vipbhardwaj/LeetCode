func generateValidStrings(n int, k int) []string {
    var res []string
    dfs(&res, n, k, "", 0)
    return res
}

func dfs(res *[]string, n, k int, s string, i int) {
    if i > k { return }
    if len(s) >= 2 && s[len(s)-2:] == "11" { return }
    if len(s) == n  {
        *res = append(*res, s)
        return
    }

    dfs(res, n, k, s+"0", i)
    dfs(res, n, k, s+"1", i + len(s))
}