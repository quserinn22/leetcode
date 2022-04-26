func climbStairs(n int) int {
    cache := make([]int, n+1)
    var dfs func(int) int
    dfs = func(i int) int {
        if cache[i] != 0 {
            return cache[i]
        }
        
        if i == 0 || i == 1 {
            return 1
        }
        
        cache[i] = dfs(i - 1) + dfs(i - 2)
        
        return cache[i]
    }
    
    return dfs(n)
}