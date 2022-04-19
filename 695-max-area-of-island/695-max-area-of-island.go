func maxAreaOfIsland(grid [][]int) int {
    ret := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                ret = max(ret, check(grid, i, j))
                
            }
        }
    }
    return ret
}

func check(grid [][]int, x int, y int) int {
    dx := []int{1, -1, 0, 0}
    dy := []int{0, 0, 1, -1}
    
    ret := 1
    grid[x][y] = 0
    
    for i := 0; i < 4; i++ {
        nx := x + dx[i]
        ny := y + dy[i]
        if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
            if grid[nx][ny] == 1 {
                ret += check(grid, nx, ny)
            }
        }
    }
    return ret
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}