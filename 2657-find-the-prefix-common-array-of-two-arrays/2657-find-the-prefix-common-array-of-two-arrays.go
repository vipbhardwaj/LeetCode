func findThePrefixCommonArray(A []int, B []int) []int {
    seen := make(map[int]int)
    res := make([]int, len(A))
    common := 0
    for i := 0; i < len(A); i++ {
        seen[A[i]]++
        if seen[A[i]] == 2 {
            common++
        }
        seen[B[i]]++
        if seen[B[i]] == 2 {
            common++
        }
        res[i] = common
    }
    return res
}