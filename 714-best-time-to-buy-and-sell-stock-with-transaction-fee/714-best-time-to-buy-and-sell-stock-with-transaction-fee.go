func dp(prices []int, cache [][]int, idx int, cnt int, fee int) int {
    if cache[idx][cnt] != -1 {
        return cache[idx][cnt]
    }
    
    if idx >= len(prices) - 1 {
        if cnt == 1 {
            return prices[idx] - fee
        } else {
            return 0
        }
    }
    
    ret := 0
    if cnt == 1 {
        ret = max(ret, dp(prices, cache, idx + 1, 0, fee) + prices[idx] - fee)
    } else {
        ret = max(ret, dp(prices, cache, idx + 1, 1, fee) -  prices[idx])
    }
    
    ret = max(ret, dp(prices, cache, idx + 1, cnt, fee))

    cache[idx][cnt] = ret
    
    return ret
}


func maxProfit(prices []int, fee int) int {
    cache := make([][]int, len(prices))
    for i := range prices {
        cache[i] = []int{-1, -1}
    }
    
    return dp(prices, cache, 0, 0, fee)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}