func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
    var have = make(map[string]bool)
    for _, s := range supplies {
        have[s] = true
    }

    var res []string
    var prevLen int
    for prevLen != len(have) {
        prevLen = len(have)
        for i, r := range recipes {
            var flag = true
            for _, ing := range ingredients[i] {
                if !have[ing] {
                    flag = false
                    break
                }
            }
            if flag && !have[r] {
                res = append(res, r)
                have[r] = true
            }
        }
    }
    return res
}