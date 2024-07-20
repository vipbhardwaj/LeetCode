func restoreMatrix(rowSum []int, colSum []int) [][]int {
    res := make([][]int, len(rowSum))
    for i := range res {
        res[i] = make([]int, len(colSum))
    }
    
    for i, j := 0, 0; i<len(rowSum) && j<len(colSum); {
        res[i][j] = min(colSum[j], rowSum[i])
        colSum[j] -= res[i][j]
        rowSum[i] -= res[i][j]
        
        if (rowSum[i] == 0) {
            i++
        }
        if (0 == colSum[j]) {
            j++
        } 
    }
    return res
}