func numMagicSquaresInside(grid [][]int) int {
    if len(grid) < 3 || len(grid[0]) < 3 {
        return 0
    }
    var res int
    for i:=0; i+3 <= len(grid); i++ {
        for j:=0; j+3 <= len(grid[0]); j++ {
            if seen(grid, i, j) {
                continue
            }
            diags := crossSum(grid, i, j)
            if diags == -1 || linearSum(grid, i, j) != diags{
                continue
            }
            res++
        }
    }
    return res
}

func seen(g [][]int, i int, j int) bool {
    seen := make([]bool, 10)
    for rc:=0; rc < 3; rc++ {
        for cc:=0; cc < 3; cc++ {
            if g[i+rc][j+cc] > 9 || seen[g[i+rc][j+cc]] || g[i+rc][j+cc] == 0 {
                return true
            } else {
                seen[g[i+rc][j+cc]] = true
            }
        }
    }
    return false
}

func linearSum(g [][]int, i int, j int) int {
    var rsum, csum int
    for sums:=0; sums < 3; sums++ {
        var currRsum, currCsum int
        for c:=0; c < 3; c++ {
            currRsum += g[i+c][j+sums]
            currCsum += g[i+sums][j+c]
        }
        if currRsum != currCsum {
            return -1
        }
        if rsum == 0 {
            rsum = currRsum
            csum = currCsum
        } else if currRsum != rsum || currCsum != csum {
            return -1
        }
    }
    return rsum
}

func crossSum(g [][]int, i int, j int) int {
    var sum1, sum2 int
    for c:=0; c < 3; c++ {
        sum1 += g[i+c][j+c]
    }
    j += 2
    for c:=0; c < 3; c++ {
        sum2 += g[i+c][j-c]
    }
    if sum1 != sum2 {
        return -1
    }
    return sum1
}