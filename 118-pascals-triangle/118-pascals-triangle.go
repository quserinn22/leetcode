func generate(numRows int) [][]int {
    ret := make([][]int, numRows + 1)
    for i := 1; i <= numRows; i++{
        ret[i] = make([]int, i)
    }
    
    for i := 1; i <= numRows; i++ {
        if i == 1 {
            ret[i][0] = 1
            continue
        }
        
        ret[i][0] = 1
        ret[i][i-1] = 1
        for j := 1; j < i/2; j++ {
            ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
            ret[i][i-j-1] = ret[i-1][j-1] + ret[i-1][j]
        }
        if i%2 != 0 {
            ret[i][i/2] = ret[i-1][(i-1)/2 - 1] + ret[i-1][(i-1)/2]
        }
    }
    
    return ret[1:]
}