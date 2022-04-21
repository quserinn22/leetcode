func orangesRotting(grid [][]int) int {
    orangeCnt := 0
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] == 1 {
                orangeCnt += 1
            }
        }
    }
    
    if orangeCnt == 0 {
        return 0
    }
    
    cnt := 0
    var done, success bool
    for {
        cnt += 1
        grid, done, success = solve(grid)
        if done {
            if success {
                return cnt
            } else {
                return -1
            }
        }
    }
}

func solve(grid [][]int) ([][]int, bool, bool) {
    dx := []int{0, 0, 1, -1}
    dy := []int{1, -1, 0, 0}
    
    cnt := 0
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] == 2 {
                for k := 0; k < 4; k++ {
                    nx := i + dx[k]
                    ny := j + dy[k]
                    if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
                        if grid[nx][ny] == 1 {
                            cnt += 1
                            grid[nx][ny] = 3       
                        }
                    }
                }
            }
        }
    }
    
    orangeCnt := 0
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] == 1 {
                orangeCnt += 1
            } else if grid[i][j] == 3 {
                grid[i][j] = 2
            }
        }
    }
    
    if orangeCnt == 0 {
        return grid, true, true
    }
        
    if cnt == 0 {
        return grid, true, false
    }
    
    return grid, false, false
}