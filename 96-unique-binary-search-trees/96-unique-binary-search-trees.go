func dp(n int, cache []int) int {
    if n == 1 || n == 0 {
        return 1
    }
    if cache[n] != 0 {
        return cache[n]
    }
    ret := 0
    for i := n-1; i >= 0; i-- {
        ret += dp(i, cache) * dp(n-i-1, cache)
    }
    
    cache[n] = ret
    return ret
}


func numTrees(n int) int {
    cache := make([]int, 20)
    
    return dp(n, cache)
}