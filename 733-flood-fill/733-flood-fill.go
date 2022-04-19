func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    image = solve(image, sr, sc, image[sr][sc], -1)   
    for i := 0; i < len(image); i++ {
        for j := 0; j < len(image[0]); j++ {
            if image[i][j] == -1 {
                image[i][j] = newColor
            }
        }
    }
    return image
}

func solve(image [][]int, sr int, sc int, origColor int, newColor int) [][]int {
    dx := []int{0, 0, 1, -1}
    dy := []int{1, -1, 0, 0}
    image[sr][sc] = newColor
    for i := 0; i < 4; i++ {
        nx := sr + dx[i]
        ny := sc + dy[i]
        if nx >= 0 && nx < len(image) && ny >= 0 && ny < len(image[0]) {
            if image[nx][ny] == origColor {
                image = solve(image, nx, ny, origColor, newColor)
            }
        }
    }
    return image
}