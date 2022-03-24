func dp(n int, cache []int) int {
    if cache[n] != 0 {
        return cache[n]
    }
    if n == 2 {
        return 1
    }
    if n == 1 {
        return 0
    }
    
    ret := n-1
    for i := 1; i < n; i++ {
        ret = max(ret, i * dp(n-i, cache))
        ret = max(ret, i * (n-i))
    }
    cache[n] = ret
    return ret
}


func integerBreak(n int) int {
    cache := make([]int, n+1)
    return dp(n, cache)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}