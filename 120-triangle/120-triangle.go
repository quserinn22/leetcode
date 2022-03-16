func dp(triangle [][]int, cache [][]int, x int, y int) int {
    if x == 0 {
        return triangle[x][y]
    }
    if cache[x][y] != 10001 * 200 {
        return cache[x][y]
    }
    
    var ret int
    if y == 0 {
        ret = dp(triangle, cache, x-1, y)
    } else if y == x {
        ret = dp(triangle, cache, x-1, y-1)
    } else {
        ret = min(dp(triangle, cache, x-1, y), dp(triangle, cache, x-1, y-1))
    }
    ret += triangle[x][y]
    
    cache[x][y] = ret
    return ret
}


func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    
    ret := 10001 * n
    cache := make([][]int, n)
    for i := range cache {
        cache[i] = make([]int, i+1)
        for j := range cache[i] {
            cache[i][j] = 10001 * 200
        }
    }
    
    for i := range triangle {
        ret = min(ret, dp(triangle, cache, n-1, i))
    }
    return ret
}


func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}