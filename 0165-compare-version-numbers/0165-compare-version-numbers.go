func compareVersion(version1 string, version2 string) int {
    var s1, s2 []string = strings.Split(version1, "."), strings.Split(version2, ".")
    var v1, v2 []int = convert(s1), convert(s2)
    var minl int = min(len(v1), len(v2))
    for i := 0; i < minl; i++ {
        if v1[i] < v2[i] {
            return -1
        } else if v1[i] > v2[i] {
            return 1
        }
    }
    if len(v1) > len(v2) {
        return 1
    } else if len(v1) < len(v2) {
        return -1
    }
    return 0
}

func convert(s []string) []int {
    var v []int
    for _, subs := range s {
        var r int
        for _, c := range subs {
            r *= 10
            r += int(c - '0')
        }
        // if r == 0 {
        //     continue
        // }
        v = append(v, r)
    }
    var i int = len(v)-1
    for; i >= 0 && v[i] == 0; i-- {}
    return v[:i+1]
}