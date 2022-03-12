func trap(height []int) int {
    cache := make([]int, 100001)
    
    maxH := 0
    sum := 0
    for _, h := range height {
        if h >= maxH {
            for i := range cache{
                sum += cache[i]
                cache[i] = 0
            }
            maxH = h
        } else {
            for i := range cache {
                if i > h {
                    if i <= maxH {
                        cache[i] += 1
                    }
                } else {
                    sum += cache[i]
                    cache[i] = 0
                }
            }
            
        }
    }
    return sum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}