func containsCycle(grid [][]byte) bool {
    visited := make(map[[2]int]bool)
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if !visited[[2]int{i, j}] {
                if dfs(grid, visited, i, j, -1, -1, grid[i][j]) {
                    return true
                }
            }
        }
    }
    return false
}

func dfs(g [][]byte, v map[[2]int]bool, i, j, pi, pj int, c byte) bool {
    if i < 0 || j < 0 || i == len(g) || j == len(g[0]) || g[i][j] != c {
        return false
    }
    if v[[2]int{i, j}] {
        return true
    }
    v[[2]int{i, j}] = true
    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for _, d := range directions {
        ni, nj := i + d[0], j + d[1]
        if ni == pi && nj == pj {
            continue
        }
        if dfs(g, v, ni, nj, i, j, c) {
            return true
        }
    }
    return false
}