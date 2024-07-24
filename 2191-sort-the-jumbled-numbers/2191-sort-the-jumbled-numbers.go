type Mapped struct {
    og int
    fake int
}

func sortJumbled(mapping []int, nums []int) []int {
    var m []Mapped
    
    for _, i := range nums {
        var n int
        var mn Mapped
        mn.og = i
        if i == 0 {
            mn.fake = mapping[0]
            m = append(m, mn)
            continue
        }
        place := 1
        for; i > 0; i /= 10 {
            n = place * mapping[i % 10] + n
            place *= 10
        }
        mn.fake = n
        m = append(m, mn)
    }
    
    sort.SliceStable(m, func(i, j int) bool {
        return m[i].fake < m[j].fake
    })
    
    for i, n := range m {
        if i == 477 || i == 478 {
            fmt.Printf("%+v\n", n)
        }
        nums[i] = n.og
    }
    return nums
}