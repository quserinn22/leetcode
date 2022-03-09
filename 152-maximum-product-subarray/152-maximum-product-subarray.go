func maxProduct(nums []int) int {
    ret := nums[0]
    maxP := nums[0]
    minP := nums[0]
    
    for _, num := range nums[1:] {
        tmp1 := maxP * num
        tmp2 := minP * num
        maxP = max(num, max(tmp1, tmp2))
        
        minP = min(num, min(tmp1, tmp2))
        
        ret = max(ret, maxP)
    }
    
    return ret
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