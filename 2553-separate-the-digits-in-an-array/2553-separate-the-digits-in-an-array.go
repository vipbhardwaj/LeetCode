func separateDigits(nums []int) []int {
    return asString(nums) // faster
    return asArray(nums) // involves another O(n) to reverse the modulated number since its back to front
}

func asString(nums []int) []int {
    var res []int
    for _, i := range nums {
        for _, j := range strconv.Itoa(i) {
            res = append(res, int(j - '0'))
        }
    }
    return res
}

func asArray(nums []int) []int {
    var res []int
    for _, i := range nums {
        var temp []int
        for i > 0 {
            temp = append(temp, i%10)
            i /= 10
        }
        reverseAppend(temp, &res)
    }
    return res
}

func reverseAppend(temp []int, res *[]int) {
    var l int = len(temp)
    for i:=l-1; i >= 0; i-- {
        *res = append(*res, temp[i])
    }
}