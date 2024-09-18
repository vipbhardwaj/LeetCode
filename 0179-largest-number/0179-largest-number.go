func largestNumber(nums []int) string {
    var numStr []string
    for _, n := range nums {
        numStr = append(numStr, strconv.Itoa(n))
    }
    sort.Slice(numStr, func(i, j int) bool {
        return numStr[i]+numStr[j] > numStr[j]+numStr[i]
    })
    if numStr[0] == "0" {
        return "0"
    }
    var res string
    for _, n := range numStr {
        res += n
    }
    return res
}