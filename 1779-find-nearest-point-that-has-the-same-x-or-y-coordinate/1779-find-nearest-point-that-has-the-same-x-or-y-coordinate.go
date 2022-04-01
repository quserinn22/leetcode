func nearestValidPoint(x int, y int, points [][]int) int {
    minIdx := -1
    minVal := 10000
    for i, point := range points {
        x1 := point[0]
        y1 := point[1]
        a := int(math.Sqrt(float64((x-x1) * (x-x1) + (y-y1) * (y-y1))))
        b := max(x, x1) - min(x, x1) + max(y, y1) - min(y, y1)
        if a == b {
            if minVal > a {
                minIdx = i
                minVal = a
            }
        }
    }
    
    return minIdx
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