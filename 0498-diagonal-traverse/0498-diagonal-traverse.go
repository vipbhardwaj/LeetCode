func findDiagonalOrder(mat [][]int) []int {
    var i, j, ic, jc, m, n int = 0, 0, -1, 1, len(mat), len(mat[0])
    var res []int
    for len(res) < m*n  {
        for i < m && j >= 0 && i >= 0 && j < n {
            res = append(res, mat[i][j])
            i += ic
            j += jc
        }
        if i < 0 && j == n {
            i, j, ic, jc = 1, n-1, 1, -1
        } else if j < 0 && i == m {
            i, j, ic, jc = m-1, 1, -1, 1
        } else if i < 0 {
            i, ic, jc = 0, 1, -1
        } else if j < 0 {
            j, ic, jc = 0, -1, 1
        } else if i == m {
            i, j, ic, jc = m-1, j+2, -1, 1
        } else if j == n {
            i, j, ic, jc = i+2, n-1, 1, -1
        }
    }
    return res
}