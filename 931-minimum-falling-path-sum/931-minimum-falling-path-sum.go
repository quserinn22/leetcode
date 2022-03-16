func dp(matrix [][]int, cache [][]int, x int, y int) int {
    if x == 0 {
        return matrix[x][y]
    }
    if cache[x][y] != 101 * len(matrix) {
        return cache[x][y]
    }
    
    ret := dp(matrix, cache, x-1, y)
    if y == 0 {
        ret = min(ret, dp(matrix, cache, x-1, y+1))
    } else if y == len(matrix)-1 {
        ret = min(ret, dp(matrix, cache, x-1, y-1))
    } else {
        ret = min(min(ret, dp(matrix, cache, x-1, y+1)), dp(matrix, cache, x-1, y-1))
    }
    
    cache[x][y] = ret + matrix[x][y]
    return ret + matrix[x][y]
}


func minFallingPathSum(matrix [][]int) int {
    n := len(matrix)
    
    ret := 101 * n
    cache := make([][]int, n)
    for i := range cache {
        cache[i] = make([]int, n)
        for j := range cache[i] {
            cache[i][j] = 101 * n
        }
    }
    
    for i := range matrix {
        ret = min(ret, dp(matrix, cache, n-1, i))
    }
    return ret
}


func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}