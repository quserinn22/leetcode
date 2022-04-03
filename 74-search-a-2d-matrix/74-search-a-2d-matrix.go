func searchMatrix(matrix [][]int, target int) bool {
    start := 1
    end := len(matrix)  * len(matrix[0])
    
    n := len(matrix[0])
    
    for start < end {
        
        mid := (end + start) / 2
        
        r, c := idxToPoint(mid, n)
        midVal := matrix[r][c]
        
        
        if midVal > target {
            end = mid
        } else if midVal < target {
            start = mid + 1
        } else {
            return true
        }
    }
    
    r, c := idxToPoint(start, n)
    return matrix[r][c] == target
}

func idxToPoint(idx int, div int) (int, int) {
    r := idx / div
    c := idx % div - 1
    
    if c == -1 {
        r -= 1
        c = div - 1
    }
    
    return r, c
}