func minZeroArray(nums []int, queries [][]int) int {
    // Defining the Diffence Array.
    var diff = make([]int, len(nums) + 1)
    // For some reason we do not populate the difference array in this question's solution algorithm, usually it is populated.
    // Line Sweep Algo!!!!
    var sum, k int
    for i := 0; i < len(nums); i++ {
        for sum + diff[i] < nums[i] {
            k++
            if k > len(queries) {
                return -1
            }
            l, r, v := queries[k-1][0], queries[k-1][1], queries[k-1][2]
            if r >= i {
                diff[max(l, i)] += v
                diff[r+1] -= v
            }
        }
        sum += diff[i]
    }
    return k
}