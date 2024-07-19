func luckyNumbers (matrix [][]int) []int {
    var mins, maxs, res []int
    
    for i:=0; i < len(matrix); i++ {
        min := 0
        for j:=0; j < len(matrix[i]); j++ {
            if min == 0 {
                min = matrix[i][j]
            } else if matrix[i][j] <= min {
                min = matrix[i][j]
            } else {
                continue
            }
        }
        mins = append(mins, min)
    }
    
    for i:=0; i < len(matrix[0]); i++ {
        max := 0
        for j:=0; j < len(matrix); j++ {
            if max == 0 {
                max = matrix[j][i]
            } else if matrix[j][i] >= max {
                max = matrix[j][i]
            } else {
                continue
            }
        }
        maxs = append(maxs, max)
    }
    
    for _, i := range maxs {
        for _, j := range mins {
            if i == j {
                res = append(res, i)
            }
        }
    }
    
    return res
}