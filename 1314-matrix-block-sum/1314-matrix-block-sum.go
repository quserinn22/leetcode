func matrixBlockSum(mat [][]int, k int) [][]int {
    m := len(mat)
    n := len(mat[0])
    
    temp := make([][]int, m + 1)
    for i := range temp {
        temp[i] = make([]int, n + 1)
    }
    
    for i := 0; i < m; i++{
        for j := 0; j < n; j++{
            temp[i + 1][j + 1] = temp[i+1][j] + temp[i][j+1] - temp[i][j] + mat[i][j];
        }
    }
    
    ret := make([][]int, m)
    for i := range ret {
        ret[i] = make([]int, n)
    }
    for i := 0; i < m; i++{
        for j := 0; j < n; j++ {
            r1 := max(0, i - k);
            c1 := max(0, j - k);
            r2 := min(m, i + k + 1);
            c2 := min(n, j + k + 1);
            ret[i][j] = temp[r2][c2] - temp[r2][c1] - temp[r1][c2] + temp[r1][c1];
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

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
