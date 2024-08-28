func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    var flag bool
    var dfs func(int, int)
    dfs = func(i int, j int) {
        if i < 0 || j < 0 || i >= len(grid2) || j >= len(grid2[0]) || grid2[i][j] == 0 {
            return
        }
        grid2[i][j] = 0
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
            if grid2[i][j] == 0 {
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