func diagonalSum(mat [][]int) int {
    n := len(mat)
    sum := 0
    for i := 0; i < n; i++ {
        sum += mat[i][i]
        if i != n-i-1 {
            sum += mat[i][n - i - 1]
        }
    }
    
    return sum
}