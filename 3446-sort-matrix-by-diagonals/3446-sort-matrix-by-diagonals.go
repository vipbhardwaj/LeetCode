func sortMatrix(grid [][]int) [][]int {
    var n int = len(grid)
    if n == 1 {
        return grid
    }
    for k:=0; k < n-1; k++ {
        var f, r []int
        for i, j := 0, k; i < n && j < n; i, j = i+1, j+1 {
            f = append(f, grid[i][j])
            r = append(r, grid[j][i])
        }
        sort.Ints(f)
        sort.Slice(r, func(a, b int) bool {
            return r[a] > r[b]
        })
        fmt.Println(f, r)
        for i, j := 0, k; i < n && j < n; i, j = i+1, j+1 {
            grid[i][j] = f[i]
            grid[j][i] = r[i]
        }
    }
    return grid
}