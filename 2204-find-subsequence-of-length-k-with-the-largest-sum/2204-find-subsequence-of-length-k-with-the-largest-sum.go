func maxSubsequence(nums []int, k int) []int {
    if k == len(nums) {
        return nums
    }
    var numsSorted [][]int
    for i, v := range nums {
        numsSorted = append(numsSorted, []int{i, v})
    }
    sort.Slice(numsSorted, func(i, j int) bool {
        // if numsSorted[i][0] > numsSorted[j][0] {
        //     return true
        // }
        return numsSorted[i][1] > numsSorted[j][1]
    })

    var resUnsorted [][]int
    for _, v := range numsSorted {
        if k == 0 {
            break
        }
        resUnsorted = append(resUnsorted, v)
        k--
    }
    sort.Slice(resUnsorted, func(i, j int) bool {
        return resUnsorted[i][0] < resUnsorted[j][0]
    })

    var res []int
    for _, v := range resUnsorted {
        res = append(res, v[1])
    }
    return res
}