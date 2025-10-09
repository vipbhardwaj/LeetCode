func minTime(skill []int, mana []int) int64 {
    var res []int64
    var curr, m, s int = mana[0], len(mana), len(skill)
    var prev int64
    for _, i := range skill {
        prev = int64(curr) * int64(i) + prev
        res = append(res, prev)
    }

    for j:=1; j < m; j++ {
        curr = mana[j]
        res[0] = res[0] + (int64(skill[0]) * int64(curr))
        for i:=1; i < s; i++ {
            res[i] = max(res[i], res[i-1]) + int64(curr) * int64(skill[i])
        }
        for i := s-1; i > 0; i-- {
            res[i-1] = res[i] - int64(curr) * int64(skill[i])
        }
    }
    return res[s-1]
}