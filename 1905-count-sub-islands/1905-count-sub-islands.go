func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    v := make([][]bool, len(grid2))
    for i:=0; i < len(grid2); i++ {
        v[i] = make([]bool, len(grid2[0]))
    }
    var flag bool
    var dfs func(int, int)
    dfs = func(i int, j int) {
        if i < 0 || j < 0 || i >= len(grid2) || j >= len(grid2[0]) || grid2[i][j] == 0 || v[i][j] {
            return
        }
        v[i][j] = true
        if grid1[i][j] != 1 {
            flag = false
        }
        dfs(i+1, j)
        dfs(i, j+1)
        dfs(i-1, j)
        dfs(i, j-1)
    }

    var res int
    for i:=0; i < len(grid2); i++ {
        for j:=0; j < len(grid1[0]); j++ {
            if v[i][j] || grid2[i][j] == 0 {
                continue
            }
            flag = true
            dfs(i, j)
            if flag {
                res++
            }
        }
    }
    return res
}