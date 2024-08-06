func minimumPushes(word string) int {
    keys := make([]int, 26)
    for _, c := range word {
        keys[c-97]++
    }
    sort.Ints(keys)
    
    var res int
    for i, k := range keys {
        if k == 0 {
            continue
        }
        kth_character := 26 - i
        if kth_character % 8 == 0 {
            res += (kth_character / 8) * k
        } else {
            res += (kth_character / 8 + 1) * k
        }
    }
    return res
}