
func dp(x int, y int, cache [][]int) int {
    if cache[x][y] != 0 {
        return cache[x][y]
    }
    
    ret := 0
    if x > 0 {
        ret += dp(x-1, y, cache)
    }
    if y > 0 {
        ret += dp(x, y-1, cache)
    }
    cache[x][y] = ret
    return ret
}

func uniquePaths(m int, n int) int {
    cache := make([][]int, m)
    for i := range cache {
        cache[i] = make([]int , n)
    }
    cache[0][0] = 1
    return dp(m-1, n-1, cache)
}