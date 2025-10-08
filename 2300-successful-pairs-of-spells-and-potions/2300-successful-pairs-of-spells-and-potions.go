func successfulPairs(spells []int, potions []int, success int64) []int {
    sort.Ints(potions)
    var res []int
    for _, i := range spells {
        var idx int = sort.SearchInts(potions, int(success / int64(i)))
        // if idx == len(potions) {
        //     res = append(res, 0)
        //     continue
        // } else if idx + 1 == len(potions) {
        //     res = append(res, len(potions))
        //     continue
        // }
        for idx < len(potions) && int64(int64(i) * int64(potions[idx])) < success {
            idx++
        }
        // fmt.Println(idx)
        res = append(res, len(potions) - idx)
    }
    return res
}