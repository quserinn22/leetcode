func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    minP := prices[0]
    maxP := prices[0]
    ret := 0
    
    for _, price := range prices {
        if price < minP {
            ret = max(ret, maxP - minP)
            minP = price
            maxP = price
        } else if price > maxP{
            maxP = price
        }
    }
    
    ret = max(ret, maxP - minP)
    
    return ret
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}