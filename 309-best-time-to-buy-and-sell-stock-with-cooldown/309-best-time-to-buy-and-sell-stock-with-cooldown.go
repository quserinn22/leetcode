func dp(prices []int, cache [][]int, idx int, cnt int) int {
    if cache[idx][cnt] != -1 {
        return cache[idx][cnt]
    }
    
    if idx >= len(prices) - 1 {
        if cnt == 1 {
            return prices[idx]
        } else {
            return 0
        }
    }
    
    ret := 0
    if cnt == 1 {
        if idx + 2 <= len(prices) - 1{
            ret = max(ret, dp(prices, cache, idx + 2, 0) + prices[idx])
        } else {
            ret = max(ret, prices[idx])
        }
    } else {
        ret = max(ret, dp(prices, cache, idx + 1, 1) -  prices[idx])
    }
    
    ret = max(ret, dp(prices, cache, idx + 1, cnt))

    cache[idx][cnt] = ret
    
    return ret
}


func maxProfit(prices []int) int {
    cache := make([][]int, len(prices))
    for i := range prices {
        cache[i] = make([]int, len(prices))
        for j := range prices {
            cache[i][j] = -1
        }
    }
    
    return dp(prices, cache, 0, 0)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}