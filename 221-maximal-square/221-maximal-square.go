func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    cache := make([][]int, m)
    for i := range cache {
        cache[i] = make([]int, n)
    }
    
    ret := 0
    for i := range matrix {
        for j := range matrix[i] {
            if i == 0 || j == 0{
                if string(matrix[i][j]) == "0" {
                    cache[i][j] = 0
                } else {
                    cache[i][j] = 1
                    ret = max(ret, 1)
                }
                continue
            }
            
            if string(matrix[i][j]) == "1" {
                v := min(min(cache[i-1][j], cache[i][j-1]), cache[i-1][j-1])
                cache[i][j] = v + 1
                ret = max(ret, cache[i][j])
            }
        }
    }
    return ret * ret
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}