func checkStraightLine(coordinates [][]int) bool {
    if coordinates[1][0] == coordinates[0][0] {
        for _, coordinate := range coordinates[1:] {
            if coordinate[0] != coordinates[0][0] {
                return false
            }
        }
        return true
    }
    
    
    check := float64(coordinates[1][1] - coordinates[0][1]) / float64(coordinates[1][0] - coordinates[0][0])
    
    for i := 1; i < len(coordinates) - 1; i++ {
        if coordinates[i+1][0] - coordinates[i][0] == 0 || float64(coordinates[i+1][1] - coordinates[i][1]) / float64(coordinates[i+1][0] - coordinates[i][0]) != check {
            return false
        }
    }
    
    return true
}