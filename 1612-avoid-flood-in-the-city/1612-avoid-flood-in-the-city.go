func avoidFlood(rains []int) []int {
    m := make(map[int]int)
    // first := make(map[int]int)
    var zeroIndices []int
    for i, r := range rains {
        if r == 0 {
            zeroIndices = append(zeroIndices, i)
            rains[i] = 1
            continue
        }
        if prev, ok := m[r]; ok {
            z := sort.SearchInts(zeroIndices, prev + 1)
            // for z=0; z < len(zeroIndices) && zeroIndices[z] < first[r]; z++ {}
            if z == len(zeroIndices) {
                return []int{}
            }
            rains[zeroIndices[z]] = r
            delete(m, r)
            zeroIndices = append(zeroIndices[:z], zeroIndices[z+1:]...)
        }
        m[r] = i
        rains[i] = -1
    }
    return rains
}