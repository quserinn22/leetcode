func getMaxLen(nums []int) int {
   
    firstMinusIdx := -1
    minusCnt := 0
    lastZeroIdx := -1
    ret := 0
    for i, num := range nums {
        if num < 0 {
            if firstMinusIdx == -1 {
                firstMinusIdx = i
            }
            minusCnt += 1
        } 
        
        
        if num == 0 {
            lastZeroIdx = i
            firstMinusIdx = -1
            minusCnt = 0
        } else {
            tmp := 0
            if minusCnt%2 == 0 {
                tmp = i - lastZeroIdx
            } else {
                tmp = i - firstMinusIdx
            }
            
            ret = max(ret, tmp)
        }
    }
    
    
    return ret
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}