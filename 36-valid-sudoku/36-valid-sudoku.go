func isValidSudoku(board [][]byte) bool {
    b := make([][]int, 9)
    for i := range b {
        b[i] = make([]int, 9)
    }
    
    for i, row := range board{
        for j, val := range row {
            if val == '.' {
                continue
            }
            
            intVal, _ := strconv.Atoi(string(val))
            b[i][j] = intVal
        }
    }
        
    for i := range b {
        valid := make([]int, 10)

        for j := range b[i] {
            if b[i][j] == 0 {
                continue
            }
            
            if valid[b[i][j]] == 1 {
                return false
            }
            valid[b[i][j]] = 1
        }
    }
            
    for i := range b {
        valid := make([]int, 10)
        for j := range b[i] {
            if b[j][i] == 0 {
                continue
            }
            
            if valid[b[j][i]] == 1 {
                return false
            }
            valid[b[j][i]] = 1
        }
    }
        
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            valid := make([]int, 10)
            
            for k := 0; k < 3; k++ {
                if b[i * 3][j * 3 + k] == 0 {
                    continue
                }
                if valid[b[i * 3][j * 3 + k]] == 1 {
                    return false
                }
                valid[b[i * 3][j * 3 + k]] = 1
            }
            
            for k := 0; k < 3; k++ {
                if b[i * 3 + 1][j * 3 + k] == 0 {
                    continue
                }
                
                if valid[b[i * 3 + 1][j * 3 + k]] == 1 {
                    return false
                }
                valid[b[i * 3 + 1][j * 3 + k]] = 1
            }
            
            for k := 0; k < 3; k++ {
                if b[i * 3 + 2][j * 3 + k] == 0 {
                    continue
                }
                
                if valid[b[i * 3 + 2][j * 3 + k]] == 1 {
                    return false
                }
                valid[b[i * 3 + 2][j * 3 + k]] = 1
            }
        }
    }
    return true
}