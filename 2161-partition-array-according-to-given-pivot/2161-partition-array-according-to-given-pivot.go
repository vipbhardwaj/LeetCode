func pivotArray(nums []int, pivot int) []int {
    var small, big []int
    var equal int 
    for _, n := range nums {
        if n < pivot {
            small = append(small, n)
        } else if n == pivot {
            equal++
        } else {
            big = append(big, n)
        }
    }
    for equal > 0 {
        small = append(small, pivot)
        equal--
    }
    for _, n := range big {
        small = append(small, n)
    }
    return small
}