func matrixReshape(mat [][]int, r int, c int) [][]int {
    ret := make([][]int, r)
    
    if r * c != len(mat) * len(mat[0]) {
        return mat
    }    
    
    a := 0
    b := 0
    for i := range ret {
        ret[i] = make([]int, c)
        for j := 0; j < c; j++ {
            ret[i][j] = mat[a][b]
            b += 1
            if b == len(mat[0]) {
                a += 1
                b = 0
            }
        }
    }
    
    
    return ret 
}