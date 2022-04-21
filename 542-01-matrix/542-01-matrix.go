func updateMatrix(mat [][]int) [][]int {
    
    for i := range mat {
        for j := range mat[0] {
            if mat[i][j] == 1 {
                mat[i][j] = 10000 * 2
            }
        }
    }
    
    for i := range mat {
        for j := range mat[0] {
            if mat[i][j] == 0 {
                mat = solve(mat, i, j)
            }
        }
    }
    
    return mat
}

func solve(mat [][]int, x int, y int) [][]int{
    dx := []int{1, -1, 0, 0}
    dy := []int{0, 0, 1, -1}
    
    v := mat[x][y]
    for i := 0; i < 4; i++ {
        nx := x + dx[i]
        ny := y + dy[i]
        if nx >= 0 && nx < len(mat) && ny >= 0 && ny < len(mat[0]) {
            if mat[nx][ny] > v + 1 {
                mat[nx][ny] = v + 1
                mat = solve(mat, nx, ny)
            }
        }
    }
    
    return mat
}