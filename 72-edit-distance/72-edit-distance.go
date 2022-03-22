func minDistance(word1 string, word2 string) int {
    n := len(word1)
    m := len(word2)
    
    cache := make([][]int, n+1)
    for i := range cache {
        cache[i] = make([]int, m+1)
    }
    
    for i := range cache {
        cache[i][0] = i
    }
    for j := range cache[0] {
        cache[0][j] = j
    }
    
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if word1[i-1] == word2[j-1] {
                cache[i][j] = cache[i-1][j-1]
            } else {
                cache[i][j] = min(cache[i][j-1], min(cache[i-1][j], cache[i-1][j-1])) + 1
            }
        }
    }
    
    return cache[n][m]
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
