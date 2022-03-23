func dp(coins []int, amount int, cache []int) int {
    if amount == 0 {
        return 0
    }
    if amount < 0 {
        return -1
    }
    
    if cache[amount] != 0 {
        return cache[amount]
    }
    
    ret := 10001
    for _, coin := range coins {
        if dp(coins, amount - coin, cache) != -1 {
            ret = min(ret,  dp(coins, amount - coin, cache) + 1)
        }
    }
    if ret == 10001 {
        ret = -1
    }
    
    cache[amount] = ret
    return ret
}

func coinChange(coins []int, amount int) int {
    cache := make([]int, 10001)
    return dp(coins, amount, cache)
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}