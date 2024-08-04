func rangeSum(nums []int, n int, left int, right int) int {
    var sum []int
    for i:=0; i < n; i++ {
        summer(&sum, 0, nums, i, n)
    }

    var res int
    sort.Ints(sum)
    fmt.Println(sum)
    for i := left-1; i < right; i++ {
        res += modded(sum[i])
    }
    return modded(res)
}

func summer(sum *[]int, num int, nums []int, i int, n int) {
    if i >= n {
        return
    }
    num += nums[i]
    *sum = append(*sum, num)
    summer(sum, num, nums, i+1, n)
}

func modded(modee int) int {
    return modee % (1e9 + 7)
}