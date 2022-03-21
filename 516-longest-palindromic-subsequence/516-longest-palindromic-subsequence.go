func longestPalindromeSubseq(s string) int {
    n := len(s)
    
    cache := make([][]int, n)
    for i := range cache {
        cache[i] = make([]int, n)
    }
    
    for i := n-1; i >= 0; i-- {
        cache[i][i] = 1
        for j := i+1; j < n; j++ {
            if s[i] == s[j] {
                cache[i][j] = cache[i+1][j-1] + 2
            } else {
                cache[i][j] = max(cache[i+1][j], cache[i][j-1])
            }
        }
    }
    
    return cache[0][n-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}