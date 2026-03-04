func numSpecial(mat [][]int) int {
    var res int
    rows := make([]int, len(mat))
    cols := make([]int, len(mat[0]))
    for i:=0; i < len(mat); i++ {
        for j:=0; j < len(mat[0]); j++ {
            if mat[i][j] == 1 {
                rows[i]++
                cols[j]++
            }
        }
    }
    for i:=0; i < len(mat); i++ {
        for j:=0; j < len(mat[0]); j++ {
            if mat[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
                res++
            }
        }
    }
    return res
    // for _, row := range mat {
    //     var seen int
    //     for _, i := range row {
    //         seen += i
    //     }
    //     if seen == 1 {
    //         rows++
    //     }
    // }
    // for i:=0; i < len(mat[0]); i++ {
    //     var saw int
    //     for j:=0; j < len(mat); j++ {
    //         saw += mat[i][j]
    //     }
    //     if saw == 1 {
    //         cols++
    //     }
    // } 
    // return min(rows, cols)
}