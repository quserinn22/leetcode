func dp(grid, cache [][]int, x, y int) int {
    if x == 0 && y == 0 {
        return grid[x][y]
    }
    if cache[x][y] != 0 {
        return cache[x][y]
    }
    
    ret := 100 * 400 + 1
    if x > 0 {
        ret = dp(grid, cache, x-1, y)
    }
    if y > 0 {
        ret = min(ret, dp(grid, cache, x, y-1))
    }
    ret += grid[x][y]    

    cache[x][y] = ret
    return ret
}


func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    
    cache := make([][]int, m)
    for i := range cache {
        cache[i] = make([]int, n)
    }
    ret := dp(grid, cache, m-1, n-1)
    return ret
}

func min(a, b int) int {
    if a > b { 
        return b
    }
    return a
}