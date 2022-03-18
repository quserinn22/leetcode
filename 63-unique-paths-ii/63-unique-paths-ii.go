
func dp(x int, y int, cache [][]int, mat [][]int) int {
    if mat[x][y] == 1 {
        return 0
    }
    if cache[x][y] != 0 {
        return cache[x][y]
    }
    
    ret := 0
    if x > 0 {
        ret += dp(x-1, y, cache, mat)
    }
    if y > 0 {
        ret += dp(x, y-1, cache, mat)
    }
    cache[x][y] = ret
    return ret
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    cache := make([][]int, m)
    for i := range cache {
        cache[i] = make([]int, n)
    }
    cache[0][0] = 1
    
    return dp(m-1, n-1, cache, obstacleGrid)
}