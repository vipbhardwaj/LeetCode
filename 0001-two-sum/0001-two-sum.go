func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
        m[target - nums[i]] = i
    }
    
    // for k, v := range m {
    //     fmt.Println(k, "-> ", v)
    // }
    
    for i := 0; i < len(nums); i++ {
        j := m[nums[i]]
        fmt.Println(j)
        if i != j && nums[i] + nums[j] == target {
            return []int{i, j}
        }
    }
    return []int{0, 0}
}